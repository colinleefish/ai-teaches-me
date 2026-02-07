# Learning Helm: Kubernetes Package Manager

Helm is the de facto package manager for Kubernetes. This guide focuses on **practical chart authoring, dependency management, and release lifecycle** — the things you actually need as an SRE.

## The Big Ideas

1. **Charts** - Reusable, versioned packages of Kubernetes manifests
2. **Templates** - Go templates + Sprig functions for dynamic manifests
3. **Values** - Layered configuration that separates config from logic
4. **Releases** - Installed instances of charts with revision history
5. **Repositories** - Chart distribution (HTTP, OCI registries)
6. **Hooks & Tests** - Lifecycle events and release validation
7. **Subcharts & Dependencies** - Composing complex applications

## Learning Path

### Phase 1: Fundamentals → `phase1-fundamentals/`

- [ ] Install Helm 3, verify with `helm version`
- [ ] Helm architecture: no Tiller, client-only in v3
- [ ] Core concepts: chart, release, revision, repository
- [ ] `helm repo add/update/list/remove`
- [ ] `helm search repo` and `helm search hub`
- [ ] `helm install`, `helm upgrade`, `helm rollback`, `helm uninstall`
- [ ] `helm list`, `helm status`, `helm history`

### Phase 2: Charts & Templates → `phase2-charts-templates/`

- [ ] Chart directory structure (`Chart.yaml`, `values.yaml`, `templates/`)
- [ ] `helm create` scaffold walkthrough
- [ ] Go template basics: `{{ .Values.x }}`, `{{ .Release.Name }}`, `{{ .Chart.Name }}`
- [ ] Built-in objects: `.Values`, `.Release`, `.Chart`, `.Capabilities`, `.Template`
- [ ] Control flow: `if/else`, `range`, `with`
- [ ] Template functions: `default`, `quote`, `toYaml`, `nindent`, `include`
- [ ] Named templates and `_helpers.tpl`
- [ ] `helm template` for local rendering and debugging

### Phase 3: Values & Configuration → `phase3-values/`

- [ ] `values.yaml` design patterns
- [ ] Overriding values: `--set`, `--set-string`, `--set-file`, `-f`
- [ ] Values merge order and precedence
- [ ] Schema validation with `values.schema.json` (JSON Schema)
- [ ] Environment-specific value files (`values-prod.yaml`, `values-staging.yaml`)

### Phase 4: Dependencies & Subcharts → `phase4-dependencies/`

- [ ] `Chart.yaml` `dependencies` field
- [ ] `helm dependency update` / `helm dependency build`
- [ ] Condition and tags for optional subcharts
- [ ] Passing values to subcharts (global vs scoped)
- [ ] Importing values from subcharts with `import-values`
- [ ] Library charts vs application charts

### Phase 5: Hooks, Tests & Lifecycle → `phase5-lifecycle/`

- [ ] Release lifecycle: install → upgrade → rollback → uninstall
- [ ] Hook types: `pre-install`, `post-install`, `pre-upgrade`, `post-upgrade`, `pre-delete`, `post-delete`, `pre-rollback`, `post-rollback`
- [ ] Hook weights and deletion policies
- [ ] Chart tests with `helm test`
- [ ] `helm upgrade --install` (upsert pattern)
- [ ] `--wait`, `--timeout`, `--atomic` flags

### Phase 6: Packaging & Distribution → `phase6-distribution/`

- [ ] `helm package` and chart versioning (SemVer)
- [ ] Chart repositories (index.yaml, ChartMuseum)
- [ ] OCI-based registries (`helm push`, `helm pull` with OCI)
- [ ] `helm registry login/logout`
- [ ] Signing charts with `helm package --sign` and provenance files

### Phase 7: Real-World Patterns → `phase7-patterns/`

- [ ] Helm in CI/CD pipelines (GitOps workflows)
- [ ] Helmfile for multi-release management
- [ ] Debugging: `helm template`, `helm get manifest`, `helm get values`
- [ ] Common pitfalls: resource naming, label selectors, immutable fields
- [ ] Upgrading CRDs (Helm's CRD limitations)
- [ ] Managing secrets (sealed-secrets, external-secrets, SOPS)
- [ ] Helm vs Kustomize: when to use which

## Quick Start

```bash
# Install Helm (macOS)
brew install helm

# Add a repo
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update

# Search for a chart
helm search repo bitnami/nginx

# Install a release
helm install my-nginx bitnami/nginx

# Check status
helm list
helm status my-nginx

# Clean up
helm uninstall my-nginx
```

## Resources

- [Helm Docs](https://helm.sh/docs/) - Official documentation
- [Chart Best Practices](https://helm.sh/docs/chart_best_practices/) - Official style guide
- [Artifact Hub](https://artifacthub.io/) - Public chart registry
- [Helmfile](https://github.com/helmfile/helmfile) - Declarative multi-release management
