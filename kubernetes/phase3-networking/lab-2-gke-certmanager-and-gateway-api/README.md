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
cluster-issuer-letsencrypt-staging.yaml  →  cert-manager 的 ACME 签发器（staging）
deployment.yaml                          →  应用本体
service.yaml                             →  ClusterIP Service
gateway.yaml                             →  Gateway（HTTP + HTTPS 监听器）
certificate.yaml                         →  cert-manager 证书请求
routes.yaml                              →  HTTPS 路由 + HTTP→HTTPS 301 跳转（两个 HTTPRoute）
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

### 3. 启用 Workload Identity

Workload Identity 是 GKE 的身份桥梁，让 Kubernetes SA（KSA）可以安全地扮演 GCP SA，无需导出密钥。

需要在两个层级启用：

1. **集群级别**：启用 Workload Identity 池（控制台：集群 → 安全 → Workload Identity）
2. **节点池级别**：启用 GKE Metadata Server（控制台：节点池 → 安全 → GKE Metadata Server → Enabled）

节点池更新会滚动重建节点，Pod 会被重新调度。

验证已启用：

```bash
gcloud container clusters describe starorigin-dev-std-zonal-asia-ne1-b \
  --zone=asia-northeast1-b \
  --format="value(workloadIdentityConfig.workloadPool)"
# 应输出：starorigin-485010.svc.id.goog
```

### 4. 配置 DNS-01 验证（Workload Identity + Cloud DNS）

cert-manager 需要往 Cloud DNS 写 TXT 记录来证明你拥有域名。通过 Workload Identity 授权，无需导出密钥。

```bash
# 1. 创建 GCP 服务账号
gcloud iam service-accounts create cert-manager-dns \
  --display-name="cert-manager DNS-01 solver"

# 2. 授予 DNS Admin 权限（在项目级别）
#    目标：项目的 IAM 策略
#    含义：cert-manager-dns 这个 SA 可以管理本项目的 DNS
gcloud projects add-iam-policy-binding starorigin-485010 \
  --member="serviceAccount:cert-manager-dns@starorigin-485010.iam.gserviceaccount.com" \
  --role="roles/dns.admin"

# 3. 绑定 KSA → GCP SA（在 SA 级别）
#    目标：cert-manager-dns 这个 SA 的 IAM 策略
#    含义：允许 cert-manager 命名空间下的 cert-manager KSA 扮演此 GCP SA
gcloud iam service-accounts add-iam-policy-binding \
  cert-manager-dns@starorigin-485010.iam.gserviceaccount.com \
  --role=roles/iam.workloadIdentityUser \
  --member="serviceAccount:starorigin-485010.svc.id.goog[cert-manager/cert-manager]"

# 4. 给 KSA 打注解，告诉 GKE Metadata Server 映射关系
kubectl annotate serviceaccount cert-manager -n cert-manager \
  iam.gke.io/gcp-service-account=cert-manager-dns@starorigin-485010.iam.gserviceaccount.com
```

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
- solver 选 `dns01.cloudDNS`，只需指定 `project: starorigin-485010`（Workload Identity 方式不需要 `serviceAccountSecretRef`，凭证通过 metadata server 自动获取）

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

### Step 5: Certificate

创建 `certificate.yaml`，要点：

- `apiVersion: cert-manager.io/v1`，`kind: Certificate`
- `secretName`：必须和 Gateway listener 的 `certificateRefs` 名称一致
- `issuerRef`：指向 Step 1 创建的 ClusterIssuer
- `dnsNames`：`gateway-whoami.gcp.xy678.vip`

DNS-01 验证不需要 Gateway IP，所以可以和 Gateway 同时进行。cert-manager 创建 Secret 后，Gateway 控制器才会完成 LB 配置并分配外部 IP。

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

### Step 6: DNS 记录

Certificate 签发完成后 Gateway 会分配外部 IP，添加 DNS A 记录：

```bash
# 获取 IP
kubectl get gateway -n gateway-whoami gateway-whoami-gateway -o jsonpath='{.status.addresses[0].value}'

# 创建 DNS 记录
gcloud dns record-sets create gateway-whoami.gcp.xy678.vip \
  --zone=gcp-xy678-vip --type=A --ttl=60 \
  --rrdatas=<GATEWAY_IP>
```

### Step 7: HTTPRoute（两个路由写在同一个 `routes.yaml` 中）

用 `---` 分隔多个资源，放在同一个 YAML 文件里是标准做法。

**HTTPS 路由：**
- `parentRefs` 指向 Gateway，`sectionName: https`（只绑定 HTTPS listener）
- `backendRefs` 指向 `gateway-whoami-service` port 80

**HTTP→HTTPS 跳转：**
- `parentRefs` 指向 Gateway，`sectionName: http`（只绑定 HTTP listener）
- 不需要 `backendRefs`——请求不会到达后端，直接返回 301
- 一条 rule，带 `RequestRedirect` filter：`scheme: https`，`statusCode: 301`

```bash
kubectl apply -f routes.yaml -n gateway-whoami
```

> **注意：** LB 配置生效需要 3-5 分钟（NEG 创建 + 健康检查通过 + URL map 同步），这是 GCP 基础设施延迟，不是 Gateway API 的问题。

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
gcloud iam service-accounts delete cert-manager-dns@starorigin-485010.iam.gserviceaccount.com
```

---

## 学习笔记

### cert-manager 是什么

Jetstack 创建，现为 CNCF 孵化项目。Kubernetes 没有内置的 TLS 证书自动化机制，cert-manager 填补了这个空白：自动签发、自动续期、支持 Let's Encrypt / Vault / Venafi 等多种 CA。它是 GKE `ManagedCertificate` 的厂商中立替代品。

安装后会创建大量资源（6 个 CRD、大量 RBAC 规则、3 个 Deployment + webhook），看起来很多但都是轻量级的，总共约 200MB 内存。

### cert-manager 的三个组件

- **cert-manager**：核心控制器，负责签发和续期证书
- **cert-manager-cainjector**：往 webhook / CRD 注入 CA 证书
- **cert-manager-webhook**：准入控制，验证你提交的 YAML 是否合法

### cert-manager 如何调用 Cloud DNS

ClusterIssuer 里配置 `dns01.cloudDNS` 后，cert-manager 加载 Google Cloud 客户端库。客户端库按 Application Default Credentials (ADC) 链查找凭证：

1. 检查 `GOOGLE_APPLICATION_CREDENTIALS` 环境变量
2. 检查 `serviceAccountSecretRef`（cert-manager 特有，从 K8s Secret 加载密钥）
3. 兜底：调用 metadata server `169.254.169.254`（节点 SA 或 Workload Identity）

cert-manager 本身是云无关的，是 ClusterIssuer YAML 里的 solver 配置决定了用哪个云的 SDK。除了 `cloudDNS`，还支持 `cloudflare`、`route53`、`azureDNS` 等。

### DNS-01 支持的 DNS 提供商

**内置（无需额外安装）：** Cloud DNS、Route53、AzureDNS、Cloudflare、Akamai

**通过 webhook 扩展：** Aliyun DNS、DNSPod / 腾讯云 DNS 等

### 节点 SA（Node Service Account）

每个 GKE 节点以一个 GCP 服务账号身份运行。本集群使用 Compute Engine 默认 SA（`586207593426-compute@developer.gserviceaccount.com`）。

查看方式：

```bash
gcloud container node-pools list \
  --cluster=starorigin-dev-std-zonal-asia-ne1-b \
  --zone=asia-northeast1-b \
  --format="table(name, config.serviceAccount)"
```

控制台路径：GKE → 集群 → 节点 → default-pool → Security → Service account

### 169.254.169.254 元数据服务

`169.254.169.254` 是链路本地地址（link-local），不经过 VPC 路由。数据包路径：

```
Pod → veth pair → 节点网络命名空间 → GCE 虚拟网卡（hypervisor 层）拦截 → Google 基础设施元数据服务响应
```

元数据服务不是跑在你 VM 上的进程，而是 Google 虚拟化层（Andromeda SDN）实现的。Hypervisor 看到目标是 `169.254.169.254` 就直接处理，数据包永远不会进入真实网络。

任何 Pod 都能访问这个地址——链路本地流量不受 Pod 网络隔离限制，只需穿过 veth pair 到达节点即可。

实测获取节点 SA 的 token：

```bash
kubectl run metadata-test --image=curlimages/curl --rm -it --restart=Never \
  -- curl -s -H "Metadata-Flavor: Google" \
  http://169.254.169.254/computeMetadata/v1/instance/service-accounts/default/token
# 返回 {"access_token":"ya29.xxx...","expires_in":2957,"token_type":"Bearer"}
```

### 三种授权 cert-manager 访问 Cloud DNS 的方式

| 方式 | 配置量 | 安全性 | 说明 |
|---|---|---|---|
| 节点 SA | 零配置 | 最弱 | 给节点 SA 加 `roles/dns.admin`，所有 Pod 共享权限 |
| SA 密钥导出 | 中等 | 中等 | 创建专用 GCP SA，导出密钥存入 K8s Secret，长期凭证有泄露风险 |
| **Workload Identity** | 中等 | **最强** | KSA 绑定 GCP SA，短期 token，按 Pod 隔离（本实验采用） |

### Workload Identity 工作原理

启用后，GKE 在每个节点上运行 `gke-metadata-server` DaemonSet，**拦截** `169.254.169.254` 请求：

```
未启用 Workload Identity：
  Pod → 169.254.169.254 → hypervisor 处理 → 返回节点 SA 的 token（所有 Pod 一样）

启用 Workload Identity：
  Pod → 169.254.169.254 → gke-metadata-server 拦截
    → 检查 Pod 的 KSA → 查找 KSA→GCP SA 绑定 → 返回对应 GCP SA 的短期 token
    → 如果没有绑定，请求被拒绝
```

同一个地址，不同的看门人。从 cert-manager 代码角度看完全透明——同样的 SDK 调用、同样的 metadata endpoint，变化纯粹在基础设施层。

### Workload Identity 绑定的三条命令详解

```bash
# 1. gcloud iam service-accounts create
#    创建 GCP SA 对象本身

# 2. gcloud projects add-iam-policy-binding starorigin-485010
#    修改【项目】的 IAM 策略
#    含义："cert-manager-dns 可以在这个项目里做 DNS Admin 的事"

# 3. gcloud iam service-accounts add-iam-policy-binding cert-manager-dns@...
#    修改【SA 自身】的 IAM 策略
#    含义："cert-manager KSA 可以扮演 cert-manager-dns"
#    --member 格式：serviceAccount:<项目WI池>[<命名空间>/<KSA名>]
```

子命令不同是因为**目标资源**不同：`projects` 修改项目策略，`iam service-accounts` 修改 SA 策略。GCP IAM 策略总是挂在某个资源上。

### Kubernetes Service Account (KSA)

每个 Pod 都以某个 KSA 身份运行。不指定则用命名空间里的 `default` KSA。每个命名空间的 `default` 是独立的对象，互不相关。

cert-manager 安装时在 `cert-manager` 命名空间创建了专用 KSA `cert-manager`，Deployment 通过 `spec.template.spec.serviceAccountName: cert-manager` 引用它。这是最佳实践：用专用 KSA 而非 `default`，便于精细授权。

### GCP SA 的跨项目能力

GCP SA 创建在某个项目中（邮箱里含项目 ID：`xxx@starorigin-485010.iam.gserviceaccount.com`），但可以被授予**任何项目**的权限。SA 的"家"只是管理归属，不限制它能操作什么。

### Gateway API 的路由能力（对比 Ingress）

Gateway API 的 `HTTPRoute` 远超 Ingress 的能力：

**匹配条件（不限于 path/host）：**
- Header 匹配：`headers: [{name: "x-env", value: "canary"}]`
- Query 参数：`queryParams: [{name: "debug", value: "true"}]`
- HTTP 方法：`method: GET`

**Filters（在请求到达后端前变换）：**
- `RequestRedirect`：重定向（本实验的 HTTP→HTTPS）
- `RequestHeaderModifier`：添加/删除/修改请求头
- `ResponseHeaderModifier`：修改响应头
- `URLRewrite`：重写路径或主机名（`ReplacePrefixMatch` 可实现路径前缀剥离）
- `RequestMirror`：镜像流量到另一个后端（影子测试）

**流量分割（金丝雀/蓝绿部署）：** 通过 `backendRefs` 的 `weight` 字段按比例分配流量。

**跨命名空间路由：** HTTPRoute 可以引用其他命名空间的 Service（需要 `ReferenceGrant` 授权）。

**角色分离：** 平台团队管 Gateway（端口、证书、IP），应用团队管 HTTPRoute（路径、后端）。各自有独立的 RBAC 权限。

**多协议：** 除了 `HTTPRoute`，还有 `TCPRoute`、`GRPCRoute`、`TLSRoute`。

Ingress 把这些功能都塞进 annotation（各厂商各自一套语法），Gateway API 用声明式的标准 API 实现。

### ACME 账号与 privateKeySecretRef

ClusterIssuer 里的 `privateKeySecretRef` 存放 ACME 账号私钥。首次使用时 cert-manager 自动向 Let's Encrypt 注册账号并生成密钥对，存入指定的 Secret。

这个密钥的作用是证明"我还是上次那个客户端"。没有它，Pod 重启后会重新注册账号，已签发的证书续期时需要重新做 DNS-01 验证。

注册账号和签发证书是两个独立步骤：ClusterIssuer Ready 只表示账号注册成功，不涉及任何证书。

### Let's Encrypt 速率限制

- **50 张证书 / 注册域名 / 周**（如 `xy678.vip`），按域名计，与账号无关
- **5 张重复证书 / 周**（完全相同的主机名组合）
- **300 个新订单 / 账号 / 3 小时**
- Staging 环境无实际限制，可随意测试

### Certificate 产生的 Secret

cert-manager 签发成功后创建一个标准的 `kubernetes.io/tls` 类型 Secret，包含：
- `tls.crt`：证书链（你的证书 + Let's Encrypt 中间 CA）
- `tls.key`：私钥

Gateway 控制器读取这个 Secret 并上传到 GCP LB 做 TLS 终止。续期时 cert-manager 自动更新 Secret，Gateway 控制器自动感知。

### Gateway 会等待 TLS Secret

如果 Gateway 的 `certificateRefs` 指向一个不存在的 Secret，GKE Gateway 控制器会报 `GWCER102: Secret not found`，不会分配外部 IP。Secret 出现后自动恢复。因此本实验中 Certificate 和 Gateway 同时 apply，等 cert-manager 签发完成后 Gateway 才分配 IP。

### ExternalDNS：自动化 DNS 记录

ExternalDNS（CNCF 项目）可以监听 Gateway/Ingress/Service 资源，自动在 Cloud DNS / Route53 / Cloudflare 等创建对应的 DNS 记录。配合 cert-manager，实现完全自动化：apply Gateway + HTTPRoute → ExternalDNS 创建 A 记录 → cert-manager 签发证书 → 流量通。

### Workload Identity 池的限制

- 每个 GKE 项目有且只有一个 Workload Identity 池：`<PROJECT_ID>.svc.id.goog`
- 绑定格式是 `[namespace/KSA]`，**不包含集群名**
- 因此：同项目内不同集群，如果有相同的 namespace + KSA 名，会共享同一个绑定
- 如需按集群隔离，使用不同的 KSA 名或不同的项目
