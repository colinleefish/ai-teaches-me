# Phase 1: Scheduling & Placement

How pods land on nodes and how to influence it.

## Topics

- [ ] How the scheduler works (filtering → scoring → binding)
- [ ] Node selectors and `nodeAffinity`
- ✅ Pod affinity and anti-affinity (`podAffinity`, `podAntiAffinity`)
- ✅ Topology keys (`kubernetes.io/hostname`, `topology.kubernetes.io/zone`)
- [ ] Topology spread constraints (`topologySpreadConstraints`)
- [ ] Taints and tolerations
- [ ] Priority classes and preemption
- [ ] DaemonSets — why your Autopilot nodes are so full

---

## 1.1 How the Scheduler Works

When you create a pod, the scheduler picks a node in three steps:

1. **Filtering** — eliminate nodes that can't run the pod (not enough CPU, wrong node selector, taints, anti-affinity violations)
2. **Scoring** — rank remaining nodes (prefer nodes with more free resources, prefer nodes that satisfy `preferred` affinity, spread evenly)
3. **Binding** — assign the pod to the highest-scoring node

The scheduler only runs at **schedule time**. It does not rebalance running pods — once a pod is on a node, it stays there until it's deleted or the node is drained.

---

## 1.2 Node Selectors and `nodeAffinity`

### nodeSelector (simple)

```yaml
spec:
  nodeSelector:
    cloud.google.com/gke-spot: "true"    # only schedule on spot nodes
```

Hard requirement — pod won't schedule if no node matches.

### nodeAffinity (expressive)

```yaml
spec:
  affinity:
    nodeAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
        nodeSelectorTerms:
          - matchExpressions:
              - key: topology.kubernetes.io/zone
                operator: In
                values:
                  - asia-northeast1-b
                  - asia-northeast1-c
      preferredDuringSchedulingIgnoredDuringExecution:
        - weight: 80
          preference:
            matchExpressions:
              - key: cloud.google.com/gke-spot
                operator: In
                values:
                  - "true"
```

This says: "Must be in zone b or c. Prefer spot nodes (weight 80)."

Operators: `In`, `NotIn`, `Exists`, `DoesNotExist`, `Gt`, `Lt`.

---

## ✅ 1.3 Pod Anti-Affinity

Controls whether pods should (or shouldn't) land on the same node/zone as other pods.

### `required` vs `preferred`

| Type | Behavior | Use when |
|---|---|---|
| `requiredDuringScheduling...` | Hard rule — pod stays Pending if unsatisfied | Co-location would **break** things |
| `preferredDuringScheduling...` | Soft rule — scheduler tries but won't block | You **want** spread but can tolerate co-location |

### Lesson learned: `required` deadlocks rollouts

We hit this firsthand. With `required` anti-affinity on a 2-node cluster:

```
Old pod on node-1 (required: no other traefik-whoami here)
Old pod on node-2 (required: no other traefik-whoami here)
New pod: needs a node with no traefik-whoami → no valid node → Pending forever
```

The rolling update can't proceed because:
- New pod can't schedule (both nodes occupied by old pods with matching labels)
- Old pods won't be removed until new pod is ready
- Deadlock.

**Fix:** switch to `preferred`. The new pod tolerates temporary co-location during rollout, then settles after old pods terminate.

### Production pattern

```yaml
spec:
  affinity:
    podAntiAffinity:
      preferredDuringSchedulingIgnoredDuringExecution:
        - weight: 100
          podAffinityTerm:
            labelSelector:
              matchLabels:
                app: traefik-whoami
            topologyKey: "kubernetes.io/hostname"
```

Key points:
- `weight: 100` — maximum preference (range 1–100)
- `podAffinityTerm` wrapper — required for `preferred` (not needed for `required`)
- `labelSelector` — which pods to avoid co-locating with
- `topologyKey` — defines "same place"

### `podAffinity` (attract, not repel)

The opposite — "schedule me **with** pods that match":

```yaml
affinity:
  podAffinity:
    preferredDuringSchedulingIgnoredDuringExecution:
      - weight: 100
        podAffinityTerm:
          labelSelector:
            matchLabels:
              app: redis
          topologyKey: "kubernetes.io/hostname"
```

Use case: put your app pod on the same node as its cache for low latency.

---

## ✅ 1.4 Topology Key

`topologyKey` defines what "same place" means for affinity/anti-affinity rules.

| `topologyKey` | Scope | Use case |
|---|---|---|
| `kubernetes.io/hostname` | Same node | Spread replicas across nodes |
| `topology.kubernetes.io/zone` | Same zone | Survive zone failures |
| `topology.kubernetes.io/region` | Same region | Multi-region spread |

Your dev cluster has 2 nodes in `asia-northeast1-b` (same zone), so `hostname` is the right choice — `zone` would have no effect since both nodes share the same zone.

For production HA across zones, you'd use:

```yaml
topologyKey: "topology.kubernetes.io/zone"
```

This ensures a zone outage doesn't take out all replicas.

---

## 1.5 Topology Spread Constraints

More fine-grained than anti-affinity — controls **how evenly** pods spread:

```yaml
spec:
  topologySpreadConstraints:
    - maxSkew: 1                              # max difference between zones
      topologyKey: "topology.kubernetes.io/zone"
      whenUnsatisfiable: DoNotSchedule        # or ScheduleAnyway
      labelSelector:
        matchLabels:
          app: my-app
```

If you have 3 zones and 6 replicas, `maxSkew: 1` ensures 2-2-2 distribution (not 4-1-1).

**Anti-affinity vs topology spread:**
- Anti-affinity: "don't put two of me on the same node" (binary)
- Topology spread: "keep me evenly distributed" (proportional)

Use topology spread for larger deployments where you care about balance, not just separation.

---

## 1.6 Taints and Tolerations

Taints on nodes **repel** pods. Tolerations on pods **allow** them to land on tainted nodes.

```bash
# Taint a node
kubectl taint nodes <node> dedicated=gpu:NoSchedule
```

```yaml
# Pod tolerates the taint
spec:
  tolerations:
    - key: "dedicated"
      operator: "Equal"
      value: "gpu"
      effect: "NoSchedule"
```

Effects:
- `NoSchedule` — don't schedule new pods (existing pods stay)
- `PreferNoSchedule` — try to avoid, but ok if necessary
- `NoExecute` — evict existing pods too

GKE uses taints for system nodes, spot nodes, and GPU nodes. That's why your pods don't land on system-only nodes.

---

## 1.7 Priority Classes and Preemption

When a cluster is full, higher-priority pods can **evict** lower-priority ones:

```yaml
apiVersion: scheduling.k8s.io/v1
kind: PriorityClass
metadata:
  name: high-priority
value: 1000000
globalDefault: false
preemptionPolicy: PreemptLowerPriority
---
spec:
  priorityClassName: high-priority
```

GKE built-in priority classes:
- `system-node-critical` (2B) — kubelet, kube-proxy
- `system-cluster-critical` (2B - 1000) — kube-dns, metrics-server

Your app pods default to priority 0. They'll never preempt system pods, but system pods can preempt yours during resource pressure.

---

## 1.8 DaemonSets

A DaemonSet runs exactly one pod per node (or per matching node). Used for:
- Log collectors (`fluentbit-gke`)
- Metrics agents (`gke-metrics-agent`)
- Network plugins (`node-local-dns`)
- Storage drivers (`pdcsi-node`)

On your Autopilot cluster, this is why 7 nodes exist for 2 app pods — each node runs ~5-6 DaemonSet pods for system services, consuming resources that Autopilot provisions nodes for.

DaemonSets bypass the scheduler — the DaemonSet controller handles placement directly.

---

## Quick Reference

```bash
# See node labels (what affinity can target)
kubectl get nodes --show-labels

# See taints on nodes
kubectl describe nodes | grep -A 3 Taints

# See which node a pod landed on
kubectl get pods -o wide

# Check scheduler events for a stuck pod
kubectl describe pod <name> | grep -A 10 Events

# See priority classes
kubectl get priorityclasses
```
