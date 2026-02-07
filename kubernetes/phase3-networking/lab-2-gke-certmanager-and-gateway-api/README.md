# 实验：Gateway API + cert-manager — 厂商中立的 Ingress 方案

## 目标

用 Gateway API（Kubernetes SIG 标准）+ cert-manager（CNCF 项目）实现与 Lab 1 相同的效果：HTTPS 暴露服务 + HTTP→HTTPS 跳转。

最终效果：`https://gateway-whoami.gcp.xy678.vip` 返回 whoami 响应，HTTP 自动 301 跳转到 HTTPS。

## 与 Lab 1 对比

|              | Lab 1（GKE 专有）               | Lab 2（厂商中立）                               |
| ------------ | ------------------------------- | ----------------------------------------------- |
| L7 路由      | `Ingress` + GCE 控制器          | `Gateway` + `HTTPRoute`                         |
| TLS 证书     | `ManagedCertificate`（GKE CRD） | `Certificate` + `ClusterIssuer`（cert-manager） |
| HTTPS 跳转   | `FrontendConfig`（GKE CRD）     | `HTTPRoute` RequestRedirect filter              |
| 可移植性     | 仅 GKE                          | 任何支持 Gateway API 的 K8s 集群                |
| 证书验证方式 | Google 内部自动验证             | ACME DNS-01（Let's Encrypt）                    |

核心区别：Lab 1 的 ManagedCertificate / FrontendConfig 是 GKE 独有 CRD，换平台就废了。Lab 2 的 Gateway API + cert-manager 是社区标准，换到 AWS/Azure/bare-metal 只需换 GatewayClass。

## 环境

- 集群：`starorigin-dev-std-zonal-asia-ne1-b`
- 命名空间：`gateway-whoami`
- 域名：`gateway-whoami.gcp.xy678.vip`（Cloud DNS 管理）
- 应用镜像：`traefik/whoami`，端口 2001

## 架构概览

```
资源关系：

ClusterIssuer ──> Certificate ──> Secret (TLS)
                                      │
Gateway (HTTP+HTTPS listeners) ───────┘
   │          │
   │          └── HTTPRoute (https) ──> Service ──> Pods
   │
   └── HTTPRoute (http-redirect) ──> 301 to HTTPS
```

你需要创建的资源（按部署顺序）：

```
cluster-issuer.yaml    →  cert-manager 的 ACME 签发器
deployment.yaml        →  应用本体
service.yaml           →  ClusterIP Service
gateway.yaml           →  Gateway（HTTP + HTTPS 监听器）
certificate.yaml       →  cert-manager 证书请求
https-route.yaml       →  HTTPS 流量路由
http-redirect.yaml     →  HTTP→HTTPS 301 跳转
```

---

## 前置准备

### 1. 启用 Gateway API

Standard 集群需要手动启用：

```bash
gcloud container clusters update starorigin-dev-std-zonal-asia-ne1-b \
  --zone=asia-northeast1-b \
  --gateway-api=standard
```

验证 GatewayClass 已就绪：

```bash
kubectl get gatewayclass
# 应该看到 gke-l7-global-external-managed 等
```

### 2. 安装 cert-manager

```bash
# 查最新版本：https://github.com/cert-manager/cert-manager/releases
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.17.1/cert-manager.yaml
```

等待所有 Pod 就绪：

```bash
kubectl get pods -n cert-manager -w
# cert-manager, cert-manager-cainjector, cert-manager-webhook 都要 Running
```

### 3. 配置 DNS-01 验证（Cloud DNS）

cert-manager 需要往 Cloud DNS 写 TXT 记录来证明你拥有域名。需要一个有 DNS 权限的 GCP 服务账号。

```bash
# 创建 GCP 服务账号
gcloud iam service-accounts create cert-manager-dns \
  --display-name="cert-manager DNS-01 solver"

# 授予 DNS Admin 权限
gcloud projects add-iam-policy-binding <PROJECT_ID> \
  --member="serviceAccount:cert-manager-dns@<PROJECT_ID>.iam.gserviceaccount.com" \
  --role="roles/dns.admin"

# 导出密钥
gcloud iam service-accounts keys create cert-manager-dns-key.json \
  --iam-account=cert-manager-dns@<PROJECT_ID>.iam.gserviceaccount.com

# 存入 cert-manager 命名空间的 Secret
kubectl create secret generic cert-manager-dns-credentials \
  --namespace=cert-manager \
  --from-file=key.json=cert-manager-dns-key.json

# 清理本地密钥文件
rm cert-manager-dns-key.json
```

> **生产环境** 应使用 Workload Identity 代替 SA 密钥导出，避免密钥泄露风险。

---

## 部署步骤

### Step 0: 创建命名空间

```bash
kubectl create namespace gateway-whoami
```

后续所有资源（除 ClusterIssuer 外）都部署到此命名空间。YAML 里加 `metadata.namespace: gateway-whoami`，或 apply 时加 `-n gateway-whoami`。

### Step 1: ClusterIssuer

创建 `cluster-issuer.yaml`，要点：

- `apiVersion: cert-manager.io/v1`，`kind: ClusterIssuer`
- 使用 ACME 协议，Let's Encrypt 生产地址：`https://acme-v02.api.letsencrypt.org/directory`
- `privateKeySecretRef`：cert-manager 自动创建，存 ACME 账号私钥
- solver 选 `dns01.cloudDNS`，指定 `project` 和 `serviceAccountSecretRef`（指向上一步创建的 Secret）

```bash
kubectl apply -f cluster-issuer.yaml

# 验证
kubectl get clusterissuer
# STATUS 应为 Ready: True
```

> **提示：** 可以先用 staging 地址 `https://acme-staging-v02.api.letsencrypt.org/directory` 测试，避免触发 Let's Encrypt 生产环境的速率限制。Staging 签发的证书不被浏览器信任，但流程完全一致。确认流程跑通后再换生产地址。

### Step 2: Deployment

创建 `deployment.yaml`，和 Lab 1 基本一致：

- 镜像 `traefik/whoami`，端口 2001
- 标签 `app: gateway-whoami`
- 可以复用 Lab 1 的 probe、resources、antiAffinity 配置
- 部署名改为 `gateway-whoami`

```bash
kubectl apply -f deployment.yaml -n gateway-whoami
kubectl get pods -n gateway-whoami -l app=gateway-whoami
```

### Step 3: Service

创建 `service.yaml`：

- `type: ClusterIP`
- selector: `app: gateway-whoami`
- `port: 80 → targetPort: 2001`
- 名称 `gateway-whoami-service`

```bash
kubectl apply -f service.yaml -n gateway-whoami
```

### Step 4: Gateway

创建 `gateway.yaml`，这是核心资源，要点：

- `apiVersion: gateway.networking.k8s.io/v1`，`kind: Gateway`
- `gatewayClassName: gke-l7-global-external-managed`（GKE 提供的全球外部 L7 LB）
- 两个 listener：
  - `https`：port 443, protocol HTTPS, hostname `gateway-whoami.gcp.xy678.vip`
    - `tls.mode: Terminate`
    - `tls.certificateRefs` 指向一个 Secret（名称要和下一步 Certificate 的 `secretName` 一致）
  - `http`：port 80, protocol HTTP, 同一 hostname

```bash
kubectl apply -f gateway.yaml -n gateway-whoami

# 等待外部 IP 分配（可能需要 1-2 分钟）
kubectl get gateway -n gateway-whoami gateway-whoami -w
# 看 ADDRESS 列出现 IP
```

### Step 5: DNS 记录

拿到 Gateway 的外部 IP 后，添加 DNS A 记录：

```bash
# 获取 IP
kubectl get gateway -n gateway-whoami gateway-whoami -o jsonpath='{.status.addresses[0].value}'

# 创建 DNS 记录
gcloud dns record-sets create gateway-whoami.gcp.xy678.vip \
  --zone=gcp-xy678-vip --type=A --ttl=60 \
  --rrdatas=<GATEWAY_IP>
```

### Step 6: Certificate

创建 `certificate.yaml`，要点：

- `apiVersion: cert-manager.io/v1`，`kind: Certificate`
- `secretName`：必须和 Gateway listener 的 `certificateRefs` 名称一致
- `issuerRef`：指向 Step 1 创建的 ClusterIssuer
- `dnsNames`：`gateway-whoami.gcp.xy678.vip`

```bash
kubectl apply -f certificate.yaml -n gateway-whoami

# 观察签发进度
kubectl get certificate -n gateway-whoami -w
# STATUS 从 False → True，通常 1-2 分钟

# 如果卡住，查看详情
kubectl describe certificate -n gateway-whoami gateway-whoami-cert
kubectl get certificaterequest -n gateway-whoami
kubectl get order -n gateway-whoami
kubectl get challenge -n gateway-whoami
```

DNS-01 验证流程：cert-manager 自动在 Cloud DNS 创建 `_acme-challenge.gateway-whoami.gcp.xy678.vip` TXT 记录 → Let's Encrypt 查询验证 → 签发证书 → cert-manager 存入 Secret → 自动清理 TXT 记录。

### Step 7: HTTPRoute（HTTPS 流量）

创建 `https-route.yaml`：

- `apiVersion: gateway.networking.k8s.io/v1`，`kind: HTTPRoute`
- `parentRefs` 指向 Gateway，`sectionName: https`（只绑定 HTTPS listener）
- `backendRefs` 指向 `gateway-whoami-service` port 80

```bash
kubectl apply -f https-route.yaml -n gateway-whoami
```

### Step 8: HTTPRoute（HTTP→HTTPS 跳转）

创建 `http-redirect.yaml`：

- `apiVersion: gateway.networking.k8s.io/v1`，`kind: HTTPRoute`
- `parentRefs` 指向 Gateway，`sectionName: http`（只绑定 HTTP listener）
- 不需要 `backendRefs`
- 一条 rule，带 `RequestRedirect` filter：`scheme: https`，`statusCode: 301`

```bash
kubectl apply -f http-redirect.yaml -n gateway-whoami
```

---

## 验证

等 Gateway 状态全部 Programmed 后（LB 配置生效可能需要几分钟）：

```bash
# 查看 Gateway 状态
kubectl get gateway -n gateway-whoami gateway-whoami
kubectl describe gateway -n gateway-whoami gateway-whoami

# 查看 HTTPRoute 状态
kubectl get httproute -n gateway-whoami

# HTTPS 测试
curl -s https://gateway-whoami.gcp.xy678.vip/

# HTTP 跳转测试
curl -sv http://gateway-whoami.gcp.xy678.vip/ 2>&1 | grep "301\|Location"

# 证书信息
curl -sv https://gateway-whoami.gcp.xy678.vip/ 2>&1 | grep "issuer\|subject"
# 应该看到 Let's Encrypt 签发
```

## 流量路径

```
用户浏览器
  → DNS: gateway-whoami.gcp.xy678.vip → <GATEWAY_IP>
  → GCP HTTP(S) Load Balancer（全球任播 IP）
  → HTTP → HTTPRoute http-redirect → 301 到 HTTPS
  → HTTPS → TLS 终止（cert-manager 签发的 Let's Encrypt 证书）
  → HTTPRoute https-route → Service → Pod gateway-whoami:2001
```

## 清理

```bash
# 删除命名空间会清理其中所有资源
kubectl delete namespace gateway-whoami

# ClusterIssuer 是集群级资源，需单独删
kubectl delete -f cluster-issuer.yaml

# 删除 DNS 记录
gcloud dns record-sets delete gateway-whoami.gcp.xy678.vip \
  --zone=gcp-xy678-vip --type=A

# （可选）卸载 cert-manager
kubectl delete -f https://github.com/cert-manager/cert-manager/releases/download/v1.17.1/cert-manager.yaml

# （可选）删除 GCP 服务账号
gcloud iam service-accounts delete cert-manager-dns@<PROJECT_ID>.iam.gserviceaccount.com
```
