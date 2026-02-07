# Phase 2: Availability & Disruptions

Keeping services up during disruptions.

## Topics

- [ ] Voluntary vs involuntary disruptions
- [ ] PodDisruptionBudgets (`minAvailable`, `maxUnavailable`)
- ✅ Rolling update strategies (`maxSurge`, `maxUnavailable`)
- ✅ Readiness vs liveness vs startup probes
- [ ] Graceful shutdown (`preStop`, `terminationGracePeriodSeconds`)
- [ ] Pod lifecycle: what happens during node drain
- [ ] Multi-zone deployments on GKE Autopilot

---

## 2.1 Voluntary vs Involuntary Disruptions

**Involuntary** — you can't prevent these:
- Node hardware failure
- Kernel panic
- VM deleted by cloud provider
- OOM kill

**Voluntary** — planned, controllable:
- `kubectl drain` (node maintenance)
- Deployment rollout (image update)
- Cluster autoscaler scaling down
- GKE node auto-upgrade

PodDisruptionBudgets only protect against **voluntary** disruptions. If a node crashes, there's nothing to negotiate with.

---

## 2.2 PodDisruptionBudgets (PDB)

A PDB tells the cluster: "don't voluntarily evict too many of my pods at once."

```yaml
apiVersion: policy/v1
kind: PodDisruptionBudget
metadata:
  name: whoami-pdb
spec:
  minAvailable: 1          # at least 1 pod must stay running
  selector:
    matchLabels:
      app: traefik-whoami
```

Or use `maxUnavailable`:

```yaml
spec:
  maxUnavailable: 1        # at most 1 pod can be down at a time
```

| Field | Meaning | Use when |
|---|---|---|
| `minAvailable: 1` | Keep at least 1 running | Small deployments |
| `minAvailable: "50%"` | Keep at least half | Larger deployments |
| `maxUnavailable: 1` | Allow at most 1 down | General purpose |

PDBs are enforced by the **eviction API** — `kubectl drain`, autoscaler, and node upgrades all respect them. Direct `kubectl delete pod` does NOT.

---

## ✅ 2.3 Rolling Update Strategy

The Deployment controller manages rollouts via ReplicaSets. When the pod template changes:

1. New ReplicaSet created with updated template
2. New RS scales up, old RS scales down (controlled by strategy)
3. Old RS kept around (for rollback)

### Default strategy: RollingUpdate

```yaml
spec:
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 25%          # how many extra pods during rollout
      maxUnavailable: 25%    # how many pods can be down during rollout
```

| Setting | Meaning | Example (4 replicas) |
|---|---|---|
| `maxSurge: 25%` | Up to 1 extra pod (5 total during rollout) | 4 old + 1 new |
| `maxUnavailable: 25%` | Up to 1 pod can be missing (3 minimum) | Kill 1 old, start 1 new |

### The deadlock we hit

With `required` pod anti-affinity on a 2-node cluster:
- maxSurge wants to create a new pod first
- New pod can't schedule (anti-affinity blocks it)
- Old pods won't be removed until new one is ready
- Deadlock → pod stuck in Pending

**Lesson:** `required` anti-affinity + rolling update + N replicas = N nodes is a trap. Use `preferred` anti-affinity or set `maxUnavailable: 1` to let the controller kill an old pod first.

### Recreate strategy

```yaml
spec:
  strategy:
    type: Recreate    # kill all old pods, then create new ones
```

Downtime guaranteed. Only use when you can't have two versions running simultaneously (e.g., database migrations, single-writer workloads).

### Useful commands

```bash
# Watch a rollout
kubectl rollout status deployment/traefik-whoami

# Rollout history
kubectl rollout history deployment/traefik-whoami

# Rollback to previous version
kubectl rollout undo deployment/traefik-whoami

# Rollback to specific revision
kubectl rollout undo deployment/traefik-whoami --to-revision=2
```

### Key insight: pod specs are immutable

Changing the Deployment updates the **template**, not existing pods. Existing pods keep their original spec (including old affinity rules) until they're replaced. This is why our stuck rollout required a delete + recreate — the old pods' `required` anti-affinity was baked into their spec.

---

## ✅ 2.4 Probes

The kubelet runs probes against each container to determine health and readiness.

### Three probe types

| Probe | Question | Failure action |
|---|---|---|
| **livenessProbe** | Is the process stuck? | Restart the container |
| **readinessProbe** | Can it handle traffic? | Remove from Service endpoints |
| **startupProbe** | Has it finished starting? | Disable liveness/readiness until pass |

### How they work together

```
Container starts
    │
    ├─ startupProbe runs (if defined)
    │   └─ Until it passes, liveness + readiness are disabled
    │
    ├─ livenessProbe starts
    │   └─ Fails → container restarted
    │
    └─ readinessProbe starts
        └─ Fails → removed from Service (no traffic)
        └─ Passes → added to Service endpoints
```

### Probe methods

```yaml
# HTTP GET (most common for web services)
livenessProbe:
  httpGet:
    path: /health
    port: 2001

# TCP socket (for non-HTTP services like databases)
livenessProbe:
  tcpSocket:
    port: 5432

# Command (run a script inside the container)
livenessProbe:
  exec:
    command:
      - cat
      - /tmp/healthy
```

### Tuning parameters

```yaml
livenessProbe:
  httpGet:
    path: /health
    port: 2001
  initialDelaySeconds: 5     # wait before first probe (default: 0)
  periodSeconds: 10           # probe every N seconds (default: 10)
  timeoutSeconds: 5           # per-probe timeout (default: 1)
  failureThreshold: 3         # failures before action (default: 3)
  successThreshold: 1         # successes to recover (default: 1)
```

### Rules and gotchas

- **`livenessProbe.successThreshold`** must be `1` — Kubernetes rejects any other value. If the process is alive once, it's alive.
- **`readinessProbe.successThreshold`** can be > 1 — useful to confirm the app is *consistently* ready before routing traffic.
- **Don't make liveness probes check dependencies** — if your DB is down, you don't want to restart your app in a loop. Liveness should only check "is my process healthy?" Readiness checks dependencies.
- **`failureThreshold: 1` is aggressive** — one slow response and your container restarts (liveness) or loses traffic (readiness). Production typically uses 3.

### What we used

```yaml
livenessProbe:
  httpGet:
    path: /health
    port: 2001
  initialDelaySeconds: 5
  periodSeconds: 10
  timeoutSeconds: 5
  failureThreshold: 3         # 3 failures = 30s before restart

readinessProbe:
  httpGet:
    path: /health
    port: 2001
  initialDelaySeconds: 5
  periodSeconds: 10
  timeoutSeconds: 5
  failureThreshold: 1
  successThreshold: 2         # 2 consecutive passes before receiving traffic
```

The traefik/whoami `/health` endpoint returns 200 by default. You can POST a status code to simulate failures:

```bash
# Make it "unhealthy" — readiness probe fails, pod removed from Service
curl -X POST -d '500' http://<pod-ip>:2001/health

# Liveness probe also fails → container restarted → health resets to 200
```

---

## 2.5 Graceful Shutdown

When a pod is terminated (rollout, drain, delete):

```
1. Pod marked as Terminating
2. Removed from Service endpoints (no new traffic)
3. preStop hook runs (if defined)
4. SIGTERM sent to container
5. Wait terminationGracePeriodSeconds (default: 30s)
6. SIGKILL if still running
```

```yaml
spec:
  terminationGracePeriodSeconds: 60    # default 30
  containers:
    - name: my-app
      lifecycle:
        preStop:
          exec:
            command: ["/bin/sh", "-c", "sleep 5"]   # drain in-flight requests
```

The `sleep 5` in preStop is a common pattern — it gives the Service endpoints time to propagate the removal before the container starts shutting down. Without it, traffic can still arrive at a container that's already received SIGTERM.

---

## 2.6 Pod Lifecycle During Node Drain

When a node is drained (`kubectl drain` or GKE auto-upgrade):

```
1. Node cordoned (no new pods scheduled)
2. For each pod on the node:
   a. Check PodDisruptionBudget — wait if budget exhausted
   b. Evict pod (same as graceful shutdown above)
   c. Pod rescheduled on another node by Deployment controller
3. Node is empty
```

DaemonSet pods are skipped by default (`--ignore-daemonsets`). Pods without a controller (bare pods) are not rescheduled — they're just gone.

---

## 2.7 Multi-Zone Deployments on GKE Autopilot

Your Autopilot cluster is regional (`asia-northeast1`), spanning 3 zones. Autopilot automatically:
- Provisions nodes across zones
- Respects topology spread if configured

For HA across zones:

```yaml
topologySpreadConstraints:
  - maxSkew: 1
    topologyKey: "topology.kubernetes.io/zone"
    whenUnsatisfiable: DoNotSchedule
    labelSelector:
      matchLabels:
        app: my-app
```

Combined with a PDB:

```yaml
apiVersion: policy/v1
kind: PodDisruptionBudget
metadata:
  name: my-app-pdb
spec:
  minAvailable: 1
  selector:
    matchLabels:
      app: my-app
```

This ensures: pods spread across zones + at least 1 always running during voluntary disruptions.

---

## Quick Reference

```bash
# Check rollout status
kubectl rollout status deployment/<name>
kubectl rollout history deployment/<name>

# Rollback
kubectl rollout undo deployment/<name>

# Check PDBs
kubectl get pdb -A

# Watch pod transitions during rollout
kubectl get pods -w -l app=<name>

# Check probe status
kubectl describe pod <name> | grep -A 5 "Liveness\|Readiness"

# Simulate drain
kubectl drain <node> --ignore-daemonsets --delete-emptydir-data --dry-run=client
```
