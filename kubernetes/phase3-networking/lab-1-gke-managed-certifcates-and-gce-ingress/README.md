# 实验：GKE GCE Ingress + ManagedCertificate 全流程

## 目标

在 GKE Standard 集群上，使用 GCE Ingress 控制器暴露一个 HTTPS 服务，配合 Google 托管证书实现自动 TLS。

最终效果：`https://traefik-whoami.gcp.xy678.vip` 返回 whoami 响应，HTTP 自动 301 跳转到 HTTPS。

## 环境

- 集群：`starorigin-dev-std-zonal-asia-ne1-b`（Zonal Standard，2 × e2-small）
- 域名：`gcp.xy678.vip`，由 Cloud DNS 管理（Cloudflare 委派子域）
- 应用镜像：`traefik/whoami`，监听端口 2001

## 资源清单与部署顺序

```
deployment.yaml      →  应用本体
service.yaml         →  ClusterIP Service
ingress.yaml         →  GCE Ingress（创建 GCP HTTP(S) LB）
managed-cert.yaml    →  Google 托管 TLS 证书
frontend-config.yaml →  HTTP→HTTPS 301 跳转
```

### 正确的部署顺序

```bash
# 1. 部署应用和 Service
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml

# 2. 先部署 Ingress，等待外部 IP 分配
kubectl apply -f ingress.yaml
kubectl get ingress traefik-whoami-ingress -w   # 等到 ADDRESS 出现

# 3. 添加 DNS 记录（指向 Ingress 外部 IP）
gcloud dns record-sets create traefik-whoami.gcp.xy678.vip \
  --zone=gcp-xy678-vip --type=A --ttl=60 \
  --rrdatas=<INGRESS_IP>

# 4. 再部署 ManagedCertificate
kubectl apply -f managed-cert.yaml
kubectl get managedcertificates -w   # 等到 STATUS 变为 Active

# 5. 最后启用 HTTP→HTTPS 跳转
kubectl apply -f frontend-config.yaml
```

## 踩过的坑

### 1. Ingress 没有分配外部 IP

**现象：** Ingress 创建后一直没有 ADDRESS，也没有 Events。

**原因：** 使用了 `spec.ingressClassName: gce`，但集群中没有注册 `IngressClass` 资源。

**解决：** 改用注解方式指定 Ingress 类：

```yaml
annotations:
  kubernetes.io/ingress.class: "gce"    # 旧方式，但更可靠
```

`spec.ingressClassName` 是新标准，但依赖 `IngressClass` 资源存在。注解方式是通用做法，所有云厂商的 Ingress 控制器都支持。

### 2. ManagedCertificate 一直卡在 Provisioning

**现象：** 证书状态 `Provisioning`，域名状态 `FailedNotVisible`。

**原因：** ManagedCertificate 在 Ingress 获得外部 IP 之前就创建了。Google 的验证系统访问不到域名，多次失败后卡死在 `FailedNotVisible` 状态，不会自动重试。

**解决：** 删除并重建 ManagedCertificate（在 Ingress IP 和 DNS 都就绪之后）：

```bash
kubectl delete managedcertificate traefik-whoami-gcp-xy678-vip-managed-cert
kubectl apply -f managed-cert.yaml
# 约 1-2 分钟后变为 Active
```

**教训：** 先有 Ingress + DNS，再建证书。证书卡住不会自愈，需要删除重建。

### 3. `required` podAntiAffinity 导致滚动更新死锁

**现象：** `kubectl rollout status` 卡住，新 Pod 一直 Pending。

**原因：** 2 个节点 + 2 个副本 + `requiredDuringScheduling` 反亲和 = 没有节点可以放新 Pod。旧 Pod 占着两个节点不让新 Pod 进来，新 Pod 起不来旧 Pod 就不会被删。

**解决：** 改用 `preferredDuringSchedulingIgnoredDuringExecution`：

```yaml
affinity:
  podAntiAffinity:
    preferredDuringSchedulingIgnoredDuringExecution:  # 软规则
      - weight: 100
        podAffinityTerm:
          labelSelector:
            matchLabels:
              app: traefik-whoami
          topologyKey: "kubernetes.io/hostname"
```

**教训：** 生产环境用 `preferred`，不用 `required`。`required` 只在共存会真正出问题时使用。

### 4. livenessProbe 的 successThreshold 只能是 1

**现象：** 尝试设置 `livenessProbe.successThreshold: 2`。

**原因：** Kubernetes 强制 liveness probe 的 successThreshold 必须为 1。逻辑是：活着就是活着，不需要"连续活两次"。readinessProbe 的 successThreshold 可以大于 1。

### 5. Pod 反亲和规则写在了 Pod spec 里，改 Deployment 不会更新已有 Pod

**现象：** 把 Deployment 的 affinity 从 `required` 改成 `preferred` 后 apply，但滚动更新仍然卡住。

**原因：** 旧 Pod 的 spec 是不可变的。改 Deployment 只影响新建的 Pod，但旧 Pod 的 `required` 规则仍然生效，阻止新 Pod 调度到同节点。

**解决：** `kubectl delete deployment` 再 `kubectl apply`，完全重建。

## 流量路径

```
用户浏览器
  → DNS: traefik-whoami.gcp.xy678.vip → 34.120.38.224
  → GCP HTTP(S) Load Balancer（全球任播 IP）
  → HTTP 请求 → FrontendConfig 301 跳转到 HTTPS
  → HTTPS 请求 → TLS 终止（ManagedCertificate）
  → URL Map 路由 → NEG（直连 Pod IP）
  → Pod traefik-whoami:2001
```

## 涉及的 GKE 特有资源

| 资源 | apiVersion | 作用 |
|---|---|---|
| ManagedCertificate | `networking.gke.io/v1` | Google 托管 TLS 证书，自动签发和续期 |
| FrontendConfig | `networking.gke.io/v1beta1` | LB 前端配置（HTTPS 跳转、SSL 策略） |
| BackendConfig | `cloud.google.com/v1` | LB 后端配置（健康检查、连接排空）— 本实验未用 |

这些都是 GKE 独有的 CRD，其他平台用 cert-manager + 各自的注解实现类似功能。

## 验证命令

```bash
# 检查所有资源状态
kubectl get deployment,svc,ingress,managedcertificates,frontendconfig

# 检查 Ingress 后端健康状态
kubectl get ingress traefik-whoami-ingress \
  -o jsonpath='{.metadata.annotations.ingress\.kubernetes\.io/backends}'

# 集群内 HTTPS 测试
kubectl run curl-test --image=curlimages/curl --rm -it --restart=Never \
  -- curl -s https://traefik-whoami.gcp.xy678.vip/

# 集群内 HTTP 跳转测试
kubectl run curl-test --image=curlimages/curl --rm -it --restart=Never \
  -- curl -sv http://traefik-whoami.gcp.xy678.vip/ 2>&1 | grep "301\|Location"
```

## 清理

```bash
kubectl delete -f ingress.yaml
kubectl delete -f frontend-config.yaml
kubectl delete -f managed-cert.yaml
kubectl delete -f service.yaml
kubectl delete -f deployment.yaml

# 删除 DNS 记录
gcloud dns record-sets delete traefik-whoami.gcp.xy678.vip \
  --zone=gcp-xy678-vip --type=A
```

注意：删除 Ingress 后 GCP 会自动清理 LB、forwarding rule、URL map 等资源，但可能需要几分钟。
