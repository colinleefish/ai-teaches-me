# Phase 3: Networking

Service discovery, traffic routing, ingress, and network policies. This is where Kubernetes gets opinionated — and where GKE adds its own layer on top.

## Your Setup Context

You have two clusters in `hs-starorigin-vpc`:

| Cluster                               | Type               | Subnet           | Pod CIDR          | Service CIDR      |
| ------------------------------------- | ------------------ | ---------------- | ----------------- | ----------------- |
| `starorigin-prod-autopilot`           | Regional Autopilot | `172.16.0.0/20`  | `10.178.128.0/17` | Autopilot-managed |
| `starorigin-dev-std-zonal-asia-ne1-b` | Zonal Standard     | `172.16.16.0/24` | `10.179.0.0/20`   | `10.179.16.0/24`  |

Key point: Pod IPs are VPC-routable (alias IPs). Service ClusterIPs are cluster-internal only (kube-proxy virtual IPs).

---

## 3.1 The Kubernetes Network Model

Four networking problems Kubernetes solves:

1. **Pod-to-pod** — every pod gets a real IP, all pods can reach each other without NAT
2. **Pod-to-service** — ClusterIP provides stable virtual IP + DNS name for a set of pods
3. **External-to-service** — LoadBalancer / Ingress exposes services outside the cluster
4. **Pod-to-external** — pods reach the internet via node's NAT (or Cloud NAT on GKE)

### Fundamental rule

> Every pod gets its own IP address. Pods can communicate with all other pods on any node without NAT.

On GKE with VPC-native mode, this is implemented via **alias IP ranges** — each node gets a `/24` slice of the pod CIDR, and GCP's VPC routing handles the rest. No overlay network needed.

---

## 3.2 Services

A Service is a stable network identity for a set of pods (selected by labels).

### ClusterIP (default)

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-app
  namespace: default
spec:
  type: ClusterIP
  selector:
    app: my-app
  ports:
    - port: 80 # the port the Service listens on
      targetPort: 8080 # the port the container listens on
      protocol: TCP
```

- Gets a virtual IP from the Service CIDR (e.g., `10.179.16.x`)
- Only reachable from within the cluster
- kube-proxy programs iptables/IPVS rules on every node to DNAT traffic to backing pods
- DNS: `my-app.default.svc.cluster.local` → ClusterIP

**How kube-proxy works:**

```
Client Pod → iptables on node → DNAT to Pod IP:8080
```

kube-proxy watches the API server for Service/Endpoints changes, then updates iptables rules on every node. The "virtual IP" never appears on any network interface — it only exists in iptables rules.

### NodePort

```yaml
spec:
  type: NodePort
  ports:
    - port: 80
      targetPort: 8080
      nodePort: 30080 # optional, auto-assigned from 30000-32767
```

- Opens a port on every node's IP
- Reachable at `<any-node-ip>:30080`
- Builds on ClusterIP (you get both)
- Rarely used directly — LoadBalancer and Ingress are better

### LoadBalancer

```yaml
spec:
  type: LoadBalancer
  ports:
    - port: 80
      targetPort: 8080
```

- On GKE: provisions a **Google Cloud Network Load Balancer** (L4)
- Gets an external IP
- Builds on NodePort (traffic: external IP → node port → pod)
- Each LoadBalancer Service = one external IP = one GCP load balancer ($$$)

### ExternalName

```yaml
spec:
  type: ExternalName
  externalName: my-db.example.com
```

- No proxying, no ClusterIP — just a CNAME DNS record
- `my-app.default.svc.cluster.local` → CNAME → `my-db.example.com`
- Use case: pointing to external databases or services

### Headless Service (ClusterIP: None)

```yaml
spec:
  clusterIP: None
  selector:
    app: my-stateful-app
```

- No virtual IP allocated
- DNS returns the pod IPs directly (A records, not a single ClusterIP)
- Used by StatefulSets for stable per-pod DNS: `pod-0.my-svc.ns.svc.cluster.local`

---

## 3.3 DNS Inside the Cluster

GKE runs `kube-dns` (based on CoreDNS). Every pod gets DNS configured to use it.

### Resolution rules

| Query                               | Resolves to                                 |
| ----------------------------------- | ------------------------------------------- |
| `my-svc`                            | ClusterIP of `my-svc` in **same namespace** |
| `my-svc.other-ns`                   | ClusterIP of `my-svc` in `other-ns`         |
| `my-svc.other-ns.svc.cluster.local` | Full FQDN — same result                     |
| `pod-0.my-svc.ns.svc.cluster.local` | Pod IP (headless service only)              |

### How it works on your cluster

```
Pod → node-local-dns (169.254.20.10) → kube-dns (ClusterIP) → upstream DNS
```

Your cluster runs `node-local-dns` as a DaemonSet — it's a local cache on each node that intercepts DNS queries before they hit the cluster `kube-dns` pods. This reduces latency and avoids conntrack race conditions.

### Debugging DNS

```bash
# From inside a pod
kubectl run dns-test --image=busybox:1.36 --rm -it --restart=Never -- nslookup kubernetes.default

# Check kube-dns pods
kubectl get pods -n kube-system -l k8s-app=kube-dns

# Check DNS config in a pod
kubectl exec <pod> -- cat /etc/resolv.conf
```

The `resolv.conf` inside pods typically looks like:

```
nameserver 169.254.20.10
search default.svc.cluster.local svc.cluster.local cluster.local
options ndots:5
```

`ndots:5` means any name with fewer than 5 dots gets the search domains appended first. So `my-svc` tries `my-svc.default.svc.cluster.local` before going external. This is why internal service names "just work" — but also why external DNS can be slow (multiple failed lookups before the bare name is tried).

---

## 3.4 Ingress

Ingress = L7 (HTTP/HTTPS) routing. One external IP, multiple services behind path/host rules.

### GKE default: GCE Ingress Controller

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: my-ingress
  annotations:
    kubernetes.io/ingress.class: "gce" # default on GKE
spec:
  rules:
    - host: app.example.com
      http:
        paths:
          - path: /api
            pathType: Prefix
            backend:
              service:
                name: api-svc
                port:
                  number: 80
          - path: /
            pathType: Prefix
            backend:
              service:
                name: frontend-svc
                port:
                  number: 80
```

What GKE creates behind the scenes:

1. **Google Cloud HTTP(S) Load Balancer** (external, global)
2. **Network Endpoint Groups (NEGs)** — points directly to pod IPs (skips NodePort)
3. **URL map** — routes `/api` vs `/` to different backend services
4. **Health checks** — per backend service

### NEGs (Network Endpoint Groups)

On VPC-native clusters (yours), GKE uses **container-native load balancing** by default:

```
Internet → GCP LB → NEG → Pod IP directly
```

Instead of the old path:

```
Internet → GCP LB → NodePort → iptables → Pod IP
```

NEGs are better: lower latency, better load distribution, no double-hop. They work because pod IPs are VPC-routable (alias IPs).

### Gateway API (the future)

Gateway API is replacing Ingress. More expressive, role-oriented:

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: Gateway
metadata:
  name: my-gateway
spec:
  gatewayClassName: gke-l7-global-external-managed
  listeners:
    - name: https
      port: 443
      protocol: HTTPS
---
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: my-route
spec:
  parentRefs:
    - name: my-gateway
  hostnames:
    - "app.example.com"
  rules:
    - matches:
        - path:
            value: /api
      backendRefs:
        - name: api-svc
          port: 80
```

Key differences from Ingress:

- **Gateway** (infra team) vs **HTTPRoute** (app team) — separation of concerns
- Supports TCP/UDP/gRPC natively, not just HTTP
- Multiple listeners on one Gateway
- Cross-namespace routing

---

## 3.5 Network Policies

Network policies are firewall rules for pods. By default, all pods can talk to all pods — network policies restrict that.

### Default deny all ingress

```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: deny-all-ingress
  namespace: default
spec:
  podSelector: {} # applies to all pods in namespace
  policyTypes:
    - Ingress
  # no ingress rules = deny all
```

### Allow specific traffic

```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: allow-api-from-frontend
  namespace: default
spec:
  podSelector:
    matchLabels:
      app: api # apply to pods with app=api
  policyTypes:
    - Ingress
  ingress:
    - from:
        - podSelector:
            matchLabels:
              app: frontend # only from pods with app=frontend
      ports:
        - port: 8080
          protocol: TCP
```

### Important on GKE

- **Standard clusters** (your dev cluster): network policies need **Calico** or **Dataplane V2** enabled. Check:

  ```bash
  gcloud container clusters describe starorigin-dev-std-zonal-asia-ne1-b \
    --zone=asia-northeast1-b \
    --format="value(networkPolicy, networkConfig.datapathProvider)"
  ```

- **Autopilot clusters**: Dataplane V2 is always enabled, network policies work out of the box.

- Network policies are **namespaced** — you set them per namespace.

- They're **additive**: if any policy selects a pod, only explicitly allowed traffic gets through. No policy = allow all.

### Common patterns

```yaml
# Allow from same namespace only
ingress:
  - from:
    - podSelector: {}

# Allow from specific namespace
ingress:
  - from:
    - namespaceSelector:
        matchLabels:
          env: production

# Allow from a CIDR (external traffic)
ingress:
  - from:
    - ipBlock:
        cidr: 172.16.0.0/12
        except:
          - 172.16.1.0/24
```

---

## 3.6 GKE-Specific Networking

### BackendConfig (GKE custom health checks)

```yaml
apiVersion: cloud.google.com/v1
kind: BackendConfig
metadata:
  name: my-backend-config
spec:
  healthCheck:
    checkIntervalSec: 15
    port: 8080
    type: HTTP
    requestPath: /healthz
  connectionDraining:
    drainingTimeoutSec: 60
```

Attach to a Service via annotation:

```yaml
metadata:
  annotations:
    cloud.google.com/backend-config: '{"default": "my-backend-config"}'
```

### ManagedCertificate (free TLS)

```yaml
apiVersion: networking.gke.io/v1
kind: ManagedCertificate
metadata:
  name: my-cert
spec:
  domains:
    - app.example.com
```

Google provisions and auto-renews the cert. Attach to Ingress:

```yaml
metadata:
  annotations:
    networking.gke.io/managed-certificates: my-cert
```

### Cloud NAT (outbound traffic)

Pods on private nodes need Cloud NAT to reach the internet. Without it, pods can receive traffic but can't initiate outbound connections (e.g., pulling images from Docker Hub, calling external APIs).

---

## 3.7 Service Mesh — When Do You Need One?

**You probably don't.** Service meshes (Istio, Anthos Service Mesh) add:

- mTLS between all services
- Fine-grained traffic management (canary, retries, circuit breaking)
- Observability (request-level metrics, distributed tracing)

**Consider one when:**

- You need mutual TLS between services (zero-trust networking)
- You have complex traffic routing needs (A/B testing, traffic mirroring)
- You need request-level observability across many services

**Don't use one when:**

- You have < 10 services
- Network policies + Ingress cover your needs
- You don't want the operational overhead (sidecar proxies, control plane)

---

## Hands-On Exercises

### Exercise 1: Explore your cluster's networking

```bash
# Check what kube-proxy mode your cluster uses
kubectl get configmap kube-proxy -n kube-system -o yaml | grep mode

# See all services and their types
kubectl get svc -A

# Check DNS resolution
kubectl run dns-test --image=busybox:1.36 --rm -it --restart=Never -- nslookup kubernetes.default

# Look at a pod's resolv.conf
kubectl run resolv-test --image=busybox:1.36 --rm -it --restart=Never -- cat /etc/resolv.conf
```

### Exercise 2: Deploy a service and test connectivity

```bash
# Create a namespace
kubectl create namespace net-lab

# Deploy a web server
kubectl create deployment web --image=nginx:alpine --replicas=2 -n net-lab

# Expose it as ClusterIP
kubectl expose deployment web --port=80 --target-port=80 -n net-lab

# Test from another pod
kubectl run curl-test --image=curlimages/curl --rm -it --restart=Never -n net-lab -- curl -s http://web.net-lab

# Test cross-namespace resolution
kubectl run curl-test --image=curlimages/curl --rm -it --restart=Never -- curl -s http://web.net-lab.svc.cluster.local

# Clean up
kubectl delete namespace net-lab
```

### Exercise 3: Network policies

```bash
# Create namespace with two apps
kubectl create namespace netpol-lab
kubectl create deployment frontend --image=nginx:alpine -n netpol-lab
kubectl create deployment backend --image=nginx:alpine -n netpol-lab
kubectl expose deployment backend --port=80 -n netpol-lab

# Verify frontend can reach backend (should work)
kubectl exec -n netpol-lab deploy/frontend -- wget -qO- --timeout=3 http://backend

# Apply default deny
kubectl apply -f - <<EOF
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: deny-all
  namespace: netpol-lab
spec:
  podSelector: {}
  policyTypes:
    - Ingress
EOF

# Try again (should timeout — if Dataplane V2 / Calico is enabled)
kubectl exec -n netpol-lab deploy/frontend -- wget -qO- --timeout=3 http://backend

# Allow frontend → backend
kubectl apply -f - <<EOF
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: allow-frontend-to-backend
  namespace: netpol-lab
spec:
  podSelector:
    matchLabels:
      app: backend
  ingress:
    - from:
        - podSelector:
            matchLabels:
              app: frontend
      ports:
        - port: 80
EOF

# Try again (should work now)
kubectl exec -n netpol-lab deploy/frontend -- wget -qO- --timeout=3 http://backend

# Clean up
kubectl delete namespace netpol-lab
```

---

## Labs

- [x] **Lab 1: GKE GCE Ingress + ManagedCertificate** — GKE 专有方案，使用 Ingress + ManagedCertificate + FrontendConfig 实现 HTTPS 暴露和 HTTP→HTTPS 跳转。[详情](./lab-1-gke-managed-certifcates-and-gce-ingress/README.md)
- [x] **Lab 2: Gateway API + cert-manager** — 厂商中立方案，使用 Gateway API + cert-manager（Let's Encrypt DNS-01）+ Workload Identity 实现相同效果。涉及 GKE Metadata Server、ACME 协议、KSA↔GCP SA 身份绑定等深度话题。[详情](./lab-2-gke-certmanager-and-gateway-api/README.md)

---

## Key Takeaways

1. **Pod IPs are real** — on VPC-native GKE, pods are first-class network citizens
2. **Services are virtual** — ClusterIP only exists in iptables rules, not on any interface
3. **DNS is the glue** — `svc-name.namespace` resolves automatically, `ndots:5` makes it work
4. **NEGs > NodePort** — GKE's container-native LB talks directly to pods
5. **Gateway API > Ingress** — if starting fresh, use Gateway API
6. **Network policies default to allow-all** — you must explicitly deny
7. **Service mesh is overkill** for most setups — start with network policies
