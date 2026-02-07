# Learning Kubernetes: The Operational Side

You already run workloads on GKE. This guide focuses on **understanding what's happening under the hood** and **operating clusters reliably** — not "what is a pod" basics.

## The Big Ideas

1. **Scheduling** - How pods land on nodes and how to influence it
2. **Availability** - Keeping services up during disruptions
3. **Networking** - Service discovery, ingress, network policies
4. **Storage** - Persistent volumes, storage classes, CSI
5. **Security** - RBAC, service accounts, pod security, secrets
6. **Observability** - Metrics, logging, debugging running workloads
7. **Resource Management** - Requests, limits, QoS, autoscaling

## Learning Path

### Phase 1: Scheduling & Placement → `phase1-scheduling/`

- [ ] How the scheduler works (filtering → scoring → binding)
- [ ] Node selectors and `nodeAffinity`
- [ ] Pod affinity and anti-affinity (`podAffinity`, `podAntiAffinity`)
- [ ] Topology spread constraints (`topologySpreadConstraints`)
- [ ] Taints and tolerations
- [ ] Priority classes and preemption
- [ ] DaemonSets — why your Autopilot nodes are so full

### Phase 2: Availability & Disruptions → `phase2-availability/`

- [ ] Voluntary vs involuntary disruptions
- [ ] PodDisruptionBudgets (`minAvailable`, `maxUnavailable`)
- [ ] Rolling update strategies (`maxSurge`, `maxUnavailable`)
- [ ] Readiness vs liveness vs startup probes
- [ ] Graceful shutdown (`preStop`, `terminationGracePeriodSeconds`)
- [ ] Pod lifecycle: what happens during node drain
- [ ] Multi-zone deployments on GKE Autopilot

### Phase 3: Networking → `phase3-networking/`

- [ ] ClusterIP, NodePort, LoadBalancer, ExternalName services
- [ ] DNS resolution inside the cluster (`<svc>.<ns>.svc.cluster.local`)
- [ ] Ingress (GKE `gce` controller, host/path routing)
- [ ] Gateway API (platform-agnostic, rewrites, header manipulation)
- [ ] Network policies (default deny, allow rules)
- [ ] GKE-specific: NEGs, BackendConfig, ManagedCertificates
- [ ] cert-manager + Let's Encrypt (platform-agnostic TLS)
- [ ] Service mesh concepts (when you need one, when you don't)

### Phase 4: Storage → `phase4-storage/`

- [ ] PersistentVolume, PersistentVolumeClaim, StorageClass
- [ ] Access modes (RWO, ROX, RWX)
- [ ] Dynamic provisioning
- [ ] StatefulSets and stable storage
- [ ] CSI drivers (what filestore-node and pdcsi-node are doing on your cluster)
- [ ] Volume snapshots and backup

### Phase 5: Security → `phase5-security/`

- [ ] RBAC: Roles, ClusterRoles, RoleBindings
- [ ] Service accounts and token projection
- [ ] Pod security standards (restricted, baseline, privileged)
- [ ] Secrets management (Kubernetes secrets, external-secrets, sealed-secrets)
- [ ] GKE Workload Identity — mapping KSA to GSA
- [ ] Network policies as security boundaries
- [ ] Image security: admission controllers, binary authorization

### Phase 6: Resource Management & Autoscaling → `phase6-resources/`

- [ ] Requests vs limits (what they actually control)
- [ ] QoS classes: Guaranteed, Burstable, BestEffort
- [ ] LimitRanges and ResourceQuotas
- [ ] Horizontal Pod Autoscaler (HPA) — CPU, memory, custom metrics
- [ ] Vertical Pod Autoscaler (VPA)
- [ ] GKE Autopilot resource management (how it differs)
- [ ] Cost optimization: right-sizing, spot/preemptible workloads

### Phase 7: Observability & Debugging → `phase7-observability/`

- [ ] `kubectl` power moves: `debug`, `top`, `logs --previous`, `exec`
- [ ] Events and what they tell you
- [ ] Metrics pipeline: Prometheus, GMP (what `collector-*` pods are doing)
- [ ] Logging: fluentbit, Cloud Logging, structured logs
- [ ] Debugging CrashLoopBackOff, ImagePullBackOff, Pending pods
- [ ] Debugging networking: DNS, connectivity, service endpoints

### Phase 8: Workload Patterns → `phase8-patterns/`

- [ ] Deployments vs StatefulSets vs DaemonSets vs Jobs/CronJobs
- [ ] Init containers and sidecar containers
- [ ] ConfigMaps and environment configuration
- [ ] Blue-green and canary deployments
- [ ] Leader election pattern
- [ ] CRDs and operators (what they are, when to use)

## Quick Reference

```bash
# Context and cluster
kubectl config current-context
kubectl cluster-info

# Inspect
kubectl get pods -A -o wide
kubectl describe pod <name> -n <ns>
kubectl top pods -n <ns>

# Debug
kubectl logs <pod> -n <ns> --previous
kubectl debug -it <pod> -n <ns> --image=busybox
kubectl get events -n <ns> --sort-by='.lastTimestamp'

# Drain / cordon
kubectl cordon <node>
kubectl drain <node> --ignore-daemonsets --delete-emptydir-data
```

## Resources

- [Kubernetes Docs](https://kubernetes.io/docs/home/) - Official reference
- [Kubernetes the Hard Way](https://github.com/kelseyhightower/kubernetes-the-hard-way) - Deep understanding
- [Learnk8s](https://learnk8s.io/) - Practical guides with diagrams
- [GKE Docs](https://cloud.google.com/kubernetes-engine/docs) - GKE-specific behavior
