# 实验：NGINX Gateway Fabric — 自建数据面的 Gateway API 方案

## 目标

用 NGINX Gateway Fabric 替代 GKE 内置的 Gateway 控制器，实现完全自建的 L7 数据面。配合 cert-manager 实现 HTTPS + HTTP→HTTPS 跳转。

最终效果：`https://ngf-whoami.gcp.xy678.vip` 返回 whoami 响应，HTTP 自动 301 跳转到 HTTPS。

## 与前两个 Lab 对比

|              | Lab 1（GKE 专有）                | Lab 2（Gateway API + GKE LB）    | Lab 3（NGINX Gateway Fabric）     |
| ------------ | -------------------------------- | -------------------------------- | --------------------------------- |
| L7 路由      | `Ingress` + GCE 控制器           | `Gateway` + `HTTPRoute`          | `Gateway` + `HTTPRoute`           |
| 数据面       | GCP HTTP(S) LB                   | GCP HTTP(S) LB                   | **NGINX（集群内 Pod）**           |
| 云 LB 用途   | L7 路由 + TLS 终止               | L7 路由 + TLS 终止               | **仅 L4 透传（TCP passthrough）** |
| TLS 终止位置 | GCP LB 边缘                      | GCP LB 边缘                      | **NGINX Pod 内部**                |
| TLS 证书     | `ManagedCertificate`             | cert-manager                     | cert-manager                      |
| GatewayClass | `gke-l7-global-external-managed` | `gke-l7-global-external-managed` | **`nginx`**                       |
| 可移植性     | 仅 GKE                           | GKE（GatewayClass 绑定）         | **任何 K8s 集群**                 |

核心区别：Lab 2 的 Gateway 控制器仍是 GKE 管的，实际创建的是 GCP HTTP(S) LB。Lab 3 的数据面完全跑在集群内（NGINX Pod），云 LB 退化为纯粹的 L4 入口，只负责把 TCP 流量转发到 NGINX Pod。

## 架构概览

```
流量路径：

用户 → Cloud L4 LB (external IP, TCP 80/443)
     → NGINX Gateway Fabric Pod (TLS 终止 + L7 路由)
     → ClusterIP Service
     → 应用 Pod

资源关系：

GatewayClass (nginx)
  └── Gateway (listeners: HTTP 80, HTTPS 443)
        ├── HTTPRoute (https) ──> Service ──> Pods
        └── HTTPRoute (http-redirect) ──> 301 to HTTPS

ClusterIssuer ──> Certificate ──> Secret (TLS)
                                      │
Gateway (HTTPS listener) ─────────────┘
```

你需要创建的资源（按部署顺序）：

```
cluster-issuer.yaml     →  cert-manager ACME 签发器（复用已有或新建）
deployment.yaml         →  应用本体
service.yaml            →  ClusterIP Service
gateway.yaml            →  Gateway（HTTP + HTTPS 监听器，gatewayClassName: nginx）
certificate.yaml        →  cert-manager 证书请求
routes.yaml             →  HTTPS 路由 + HTTP→HTTPS 301 跳转
```

---

## 环境

- 集群：`starorigin-dev-std-zonal-asia-ne1-b`
- 命名空间：`ngf-whoami`（应用）、`nginx-gateway`（NGINX Gateway Fabric 本身）
- 域名：`ngf-whoami.gcp.xy678.vip`（Cloud DNS 管理）
- 应用镜像：`traefik/whoami`，端口 2001

---

## 前置准备

### 1. 安装 NGINX Gateway Fabric

NGF 通过 Helm 安装，包含控制面（controller）和数据面（NGINX Pod）。

**安装 Gateway API CRDs（NGF 版本）：**

集群上已有 GKE 安装的 Gateway API CRDs。NGF 需要特定版本，先检查兼容性：

```bash
# 查看当前 CRD 版本
kubectl get crd gateways.gateway.networking.k8s.io -o jsonpath='{.metadata.annotations.gateway\.networking\.k8s\.io/bundle-version}'
# 当前集群：v1.4.0（GKE 管理）
```

NGF 版本与 Gateway API 版本对应关系（[完整矩阵](https://github.com/nginx/nginx-gateway-fabric#technical-specifications)）：

| NGF     | Gateway API | Kubernetes |
| ------- | ----------- | ---------- |
| 2.4.1   | 1.4.1       | 1.25+      |
| 2.3.0   | 1.4.1       | 1.25+      |
| 2.2.x   | 1.3.0       | 1.25+      |
| 2.0.x   | 1.3.0       | 1.25+      |

当前集群 CRD 为 v1.4.0，NGF v2.4.1 期望 v1.4.1。差异是小版本补丁，大概率兼容。如果遇到问题，可以用 NGF 提供的 CRDs 覆盖：

```bash
kubectl kustomize "https://github.com/nginx/nginx-gateway-fabric/config/crd/gateway-api/standard?ref=v2.4.1" | kubectl apply -f -
```

> **注意：** 覆盖 CRDs 可能影响 GKE 内置的 GatewayClass。如果你同时使用 GKE Gateway 和 NGF，需要确认 CRD 版本双方都兼容。本实验只用 NGF，所以可以直接覆盖。如果之后想恢复 GKE 管理，重新启用 `--gateway-api=standard` 即可。

**通过 Helm 安装 NGF（开源版）：**

```bash
helm install ngf oci://ghcr.io/nginx/charts/nginx-gateway-fabric \
  --create-namespace -n nginx-gateway \
  --wait
```

验证：

```bash
# 控制面 Pod 就绪
kubectl get pods -n nginx-gateway
# 应看到 ngf-nginx-gateway-fabric-xxx Running

# GatewayClass 已注册
kubectl get gatewayclass nginx
# ACCEPTED 应为 True
```

### 2. cert-manager（已安装）

确认 cert-manager 仍在运行：

```bash
kubectl get pods -n cert-manager
# 三个 Pod 都应为 Running
```

### 3. ClusterIssuer

集群上还有 `letsencrypt-staging`。如果要用生产证书（浏览器信任），需要创建一个生产 ClusterIssuer，或者继续用 staging 测试。

检查现有状态：

```bash
kubectl get clusterissuer
```

---

## 部署步骤

### Step 0: 创建命名空间

```bash
kubectl create namespace ngf-whoami
```

### Step 1: ClusterIssuer

如果现有的 `letsencrypt-staging` 满足需求可以复用，否则创建新的。

要点和 Lab 2 一致：

- ACME + DNS-01 + Cloud DNS
- Workload Identity 授权（Lab 2 已配置，应仍然生效）

```bash
# 验证 Workload Identity 绑定仍在
kubectl get sa cert-manager -n cert-manager -o jsonpath='{.metadata.annotations.iam\.gke\.io/gcp-service-account}'
# 应返回 cert-manager-dns@starorigin-485010.iam.gserviceaccount.com
```

### Step 2: Deployment

创建 `deployment.yaml`：

- 镜像 `traefik/whoami`，端口 2001
- 标签 `app: ngf-whoami`
- 部署到 `ngf-whoami` 命名空间

```bash
kubectl apply -f deployment.yaml -n ngf-whoami
kubectl get pods -n ngf-whoami -l app=ngf-whoami
```

### Step 3: Service

创建 `service.yaml`：

- `type: ClusterIP`
- selector: `app: ngf-whoami`
- `port: 80 → targetPort: 2001`

```bash
kubectl apply -f service.yaml -n ngf-whoami
```

### Step 4: Gateway

创建 `gateway.yaml`，**这是和 Lab 2 最大的不同**：

- `gatewayClassName: nginx`（不是 `gke-l7-global-external-managed`）
- 两个 listener：
  - `https`：port 443, protocol HTTPS, hostname `ngf-whoami.gcp.xy678.vip`
    - `tls.mode: Terminate`
    - `tls.certificateRefs` 指向 TLS Secret（和 Certificate 的 `secretName` 一致）
  - `http`：port 80, protocol HTTP, 同一 hostname

```bash
kubectl apply -f gateway.yaml -n ngf-whoami
```

> **和 Lab 2 的区别：** Lab 2 的 Gateway 创建后，GKE 控制器在 GCP 上创建一个 HTTP(S) LB（L7，全球任播 IP）。Lab 3 的 Gateway 创建后，NGF 控制器在集群内创建一个 LoadBalancer Service（GCP 分配一个 Network LB，L4），指向 NGINX Pod。

```bash
# 查看 NGF 创建的 LoadBalancer Service
kubectl get svc -n nginx-gateway
# 应看到 TYPE=LoadBalancer，有 EXTERNAL-IP

# 也可以从 Gateway 状态获取地址
kubectl get gateway -n ngf-whoami -w
# 等待 ADDRESS 出现
```

### Step 5: Certificate

创建 `certificate.yaml`：

- `secretName` 和 Gateway listener 的 `certificateRefs` 一致
- `issuerRef` 指向 ClusterIssuer
- `dnsNames`: `ngf-whoami.gcp.xy678.vip`

```bash
kubectl apply -f certificate.yaml -n ngf-whoami

# 观察签发进度
kubectl get certificate -n ngf-whoami -w
```

> **注意：** NGF 的 Gateway 不像 GKE Gateway 那样会等 Secret 才分配 IP。NGF 的 IP 来自 LoadBalancer Service（和 TLS 无关），所以 Gateway 地址会立即可用。但 HTTPS listener 在 Secret 就绪前不会正常工作。

### Step 6: DNS 记录

```bash
# 获取 Gateway IP（来自 NGF 的 LoadBalancer Service）
kubectl get gateway -n ngf-whoami -o jsonpath='{.status.addresses[0].value}'

# 创建 DNS 记录
gcloud dns record-sets create ngf-whoami.gcp.xy678.vip \
  --zone=gcp-xy678-vip --type=A --ttl=60 \
  --rrdatas=<GATEWAY_IP>
```

### Step 7: HTTPRoute

创建 `routes.yaml`，两个 HTTPRoute 写在同一个文件中：

**HTTPS 路由：**

- `parentRefs` 指向 Gateway，`sectionName: https`
- `backendRefs` 指向 Service port 80

**HTTP→HTTPS 跳转：**

- `parentRefs` 指向 Gateway，`sectionName: http`
- `RequestRedirect` filter：`scheme: https`，`statusCode: 301`

```bash
kubectl apply -f routes.yaml -n ngf-whoami
```

---

## 验证

```bash
# 查看 Gateway 状态
kubectl get gateway -n ngf-whoami
kubectl describe gateway -n ngf-whoami

# 查看 HTTPRoute 状态
kubectl get httproute -n ngf-whoami

# 查看 NGF 的 LoadBalancer Service
kubectl get svc -n nginx-gateway

# HTTPS 测试
curl -s https://ngf-whoami.gcp.xy678.vip/

# HTTP 跳转测试
curl -sv http://ngf-whoami.gcp.xy678.vip/ 2>&1 | grep "301\|Location"

# 证书信息
curl -sv https://ngf-whoami.gcp.xy678.vip/ 2>&1 | grep "issuer\|subject"

# 如果用 staging 证书（不被信任），加 -k 跳过验证
curl -sk https://ngf-whoami.gcp.xy678.vip/
```

## 流量路径

```
用户浏览器
  → DNS: ngf-whoami.gcp.xy678.vip → <GATEWAY_IP>
  → GCP Network Load Balancer（L4，区域级 IP，仅 TCP 转发）
  → NGINX Gateway Fabric Pod（集群内）
  → HTTP → HTTPRoute http-redirect → 301 到 HTTPS
  → HTTPS → TLS 终止（cert-manager 签发的证书，在 NGINX 内完成）
  → HTTPRoute https-route → Service → Pod ngf-whoami:2001
```

和 Lab 2 的流量路径对比：

```
Lab 2:  用户 → GCP HTTP(S) LB (L7, TLS终止, URL路由) → NEG → Pod
Lab 3:  用户 → GCP Network LB (L4, TCP透传) → NGINX Pod (TLS终止, L7路由) → Service → Pod
```

Lab 2 的智能在云端（GCP LB 做一切），Lab 3 的智能在集群内（NGINX 做一切，云 LB 只是个门）。

## 清理

```bash
# 删除应用命名空间（清理其中所有资源）
kubectl delete namespace ngf-whoami

# 删除 DNS 记录
gcloud dns record-sets delete ngf-whoami.gcp.xy678.vip \
  --zone=gcp-xy678-vip --type=A

# 卸载 NGINX Gateway Fabric
helm uninstall ngf -n nginx-gateway
kubectl delete namespace nginx-gateway

# （可选）删除 NGF 安装的 Gateway API CRDs
# 注意：如果还要用 GKE 内置 Gateway，不要删
kubectl kustomize "https://github.com/nginx/nginx-gateway-fabric/config/crd/gateway-api/standard?ref=v2.4.1" | kubectl delete -f -
```

---

## 学习笔记

### NGINX Gateway Fabric 是什么

NGINX 官方的 Gateway API 实现（开源，Apache 2.0）。它不是 NGINX Ingress Controller（那是基于 Ingress API 的老项目）。NGF 从零设计，原生支持 Gateway API。

架构：控制面（Go 写的 controller，watch Gateway API 资源）+ 数据面（NGINX 进程，处理实际流量）。

### NGF vs GKE 内置 Gateway 控制器

|          | GKE Gateway 控制器               | NGINX Gateway Fabric                  |
| -------- | -------------------------------- | ------------------------------------- |
| 数据面   | GCP HTTP(S) LB（云基础设施）     | NGINX Pod（集群内）                   |
| TLS 终止 | GCP LB 边缘                      | NGINX Pod                             |
| 扩缩容   | GCP 自动管理                     | 你自己管（HPA / 手动）                |
| 生效速度 | 慢（GCP LB 配置 3-5 分钟）       | 快（NGINX reload 秒级）               |
| 可观测性 | Cloud Logging / Monitoring       | NGINX access log / Prometheus metrics |
| 成本     | GCP LB 按量计费                  | 仅 Pod 资源 + 一个 L4 LB              |
| 功能     | GCP 生态集成（IAP、Cloud Armor） | NGINX 生态（rate limit、snippets）    |

### 为什么还需要一个云 LB

NGINX Pod 跑在集群内，没有公网 IP。需要一个入口把外部流量引进来。LoadBalancer Service 让云厂商创建一个 L4 LB，把 TCP 80/443 流量转发到 NGINX Pod 的 NodePort。

你也可以不用云 LB——用 NodePort 直接暴露，或者 `hostNetwork: true` 让 NGINX 直接绑节点 IP。但 LoadBalancer Service 是最常见的做法，因为你得到一个稳定的外部 IP。

### Gateway CRD 版本冲突问题

GKE 启用 `--gateway-api=standard` 后会自动安装和管理 Gateway API CRDs。NGF 也需要特定版本的 CRDs。如果版本不一致：

- GKE 可能在集群升级时覆盖你装的 CRDs
- NGF 可能因为缺少某些字段而报错

最安全的做法：如果只用 NGF，可以关闭 GKE 的 Gateway API（`--gateway-api=disabled`），然后装 NGF 自己的 CRDs。如果需要同时使用两者，确保 CRD 版本兼容。

### TLS 终止位置的区别

**Lab 2（GCP LB 边缘终止）：**

- 证书上传到 GCP LB，TLS 在 Google 基础设施完成
- NGINX/后端收到的是明文 HTTP
- 优点：GCP 全球边缘网络加速 TLS 握手
- 缺点：证书管理绑定 GCP

**Lab 3（NGINX Pod 内终止）：**

- 证书存在 K8s Secret，NGINX 读取并处理 TLS
- 云 LB 只看到加密的 TCP 流量
- 优点：完全可移植，证书管理和云无关
- 缺点：TLS 握手延迟略高（流量要先到集群才终止）

### cert-manager 和 NGF 的配合

cert-manager 签发证书存入 Secret → Gateway 的 `certificateRefs` 引用该 Secret → NGF 读取 Secret 配置 NGINX 的 TLS。续期时 cert-manager 更新 Secret，NGF 检测到变化自动 reload NGINX。

和 Lab 2 的 cert-manager 用法几乎一样，唯一的区别是消费 Secret 的对象：Lab 2 是 GKE Gateway 控制器（上传到 GCP LB），Lab 3 是 NGF（加载到 NGINX 进程）。
