# Lab 1: Pod 调度实验

集群：`starorigin-prod-std-reg-asia-ne1-maxin`（Standard / Regional / asia-northeast1）
节点池：`general-05c2g`（e2-small × 3，a/b/c 三个可用区各一个节点）

## 实验 0：无调度策略

replicas=3，无 affinity。

**结果**：Pod 均匀分布，每个 Node 一个。

## 实验 1：节点亲和性和反亲和性

## 实验 1.1：podAffinity preferred 同节点

### 实验方案

replicas=3，`podAffinity.preferredDuringSchedulingIgnoredDuringExecution`，weight=100，topologyKey=`kubernetes.io/hostname`。

```yaml
replicas: 3

...

spec:
    affinity:
    podAffinity:
        preferredDuringSchedulingIgnoredDuringExecution:
        - weight: 100
            podAffinityTerm:
            labelSelector:
                matchLabels:
                app: whoami
            topologyKey: "kubernetes.io/hostname"
```

**结果**：Pod 仍然分散到 3 个节点。可能原因：第一个 Pod 调度时没有匹配的 Pod 存在，affinity 无信号；加上 SelectorSpread 插件会主动把同一 ReplicaSet 的 Pod 分散到不同节点。

## 实验 1.2：podAffinity preferred 同节点，先 1 replica 再扩到 3

同样的 affinity 配置。先 replicas=1 等 Running，再 `kubectl scale --replicas=3`。

**结果**：3 个 Pod 全部调度到同一个节点。说明 preferred podAffinity 本身是生效的，实验 1.1 分散的原因是 3 个 Pod 同时创建，第一个 Pod 调度时无匹配 Pod，affinity 无信号。

## 实验 1.3：podAffinity required 同节点

### 实验方案

先把 Pod 全删掉，然后 apply 这个方案：

- replicas=3
- `podAffinity.requiredDuringSchedulingIgnoredDuringExecution`
- matchLabels, app: whoami
- topologyKey=`kubernetes.io/hostname`

```yaml
spec:
  affinity:
  podAffinity:
    requiredDuringSchedulingIgnoredDuringExecution:
      - labelSelector:
          matchLabels:
            app: whoami
          topologyKey: "kubernetes.io/hostname"
```

**结果**：3 个 Pod 能够正常创建并且全部调度到同一个节点。这里可能会被误认为无法调度，但其实确实调度上了。

## 实验 1.4：podAffinity required，replica 数量超过可承载范围

先把 Pod 全删掉，然后调整 resources 使得每个 Node 能容纳 2 个这样的 Pod，然后设置 replica 数为 4。

**结果**：2 个 Pod 被调度上了，另外 2 个 Pending。

调度过程分析（4 个 Pod 同时创建，3 个 Node）：

1. **拓荒阶段**：调度器处理 Pod-1，全集群没有 `app: whoami` 的 Pod。根据自我亲和规则，第一个 Pod 可以在任意满足资源条件的节点落户。Pod-1 → Node A。
2. **锚定阶段**：调度器处理 Pod-2，发现 Node A 已有匹配 Pod，Node B/C 没有，过滤阶段直接剔除 B/C。Node A 还有空位，Pod-2 → Node A。此时 Node A 资源耗尽。
3. **锁死阶段**：调度器处理 Pod-3、Pod-4。Node A 是唯一满足 affinity 的节点但资源不足，Node B/C 不满足 affinity 规则。结果：Pending。

## 实验 1.5：硬亲和，2 个 Running 在同一个节点上，另外 2 个 Pending，然后做 Node Eviction

**结果**：会被重新调度，开荒策略重新生效

## 实验 1.6：硬亲和，但是 require 了一个不存在的 App Label

```yaml
spec:
  affinity:
  podAffinity:
    requiredDuringSchedulingIgnoredDuringExecution:
      - labelSelector:
          matchLabels:
            app: whoami
          topologyKey: "kubernetes.io/hostname"
```

**结果**：全 Pending 了。

