# Exercise 5: Advanced Features

## Objective
Learn advanced UV features: dependency resolution, compilation, workspaces, and scripts.

---

## Tasks

### Task 5.1: Compile Requirements
Compile dependencies with full resolution.

```bash
cd uv/demo-project/myapp
uv pip compile pyproject.toml -o requirements-compiled.txt
```

**Expected Result:**
- `requirements-compiled.txt` created
- Contains all dependencies with exact versions
- Includes transitive dependencies
- Platform-specific markers if needed

**Verify:**
```bash
cat requirements-compiled.txt | head -20
wc -l requirements-compiled.txt
```

---

### Task 5.2: Upgrade All Dependencies
Update all dependencies to latest compatible versions.

```bash
uv lock --upgrade
```

**Expected Result:**
- `uv.lock` updated with latest versions
- Respects version constraints in `pyproject.toml`
- Shows which packages were upgraded

**Verify:**
```bash
git diff uv.lock  # See what changed
uv sync  # Install upgraded versions
```

---

### Task 5.3: Upgrade Specific Package
Update only one package.

```bash
uv lock --upgrade-package fastapi
```

**Expected Result:**
- Only `fastapi` (and its deps) updated
- Other packages unchanged
- Minimal lockfile changes

**Verify:**
```bash
git diff uv.lock | grep fastapi
```

---

### Task 5.4: Create Project with Scripts
Define custom scripts in `pyproject.toml`.

Add to `pyproject.toml`:
```toml
[project.scripts]
hello = "myapp.main:greet"
dev = "myapp.cli:dev_command"
```

Create corresponding functions in your code.

**Expected Result:**
- Scripts available as commands after install
- Can run with `uv run hello`

**Verify:**
```bash
uv sync
uv run hello
```

---

### Task 5.5: Use Environment Variables
Configure UV behavior with environment variables.

```bash
export UV_CACHE_DIR=/tmp/uv-cache
export UV_NO_CACHE=1
uv pip install requests
```

**Expected Result:**
- UV uses custom cache location
- Or disables cache entirely
- Affects download and resolution

**Verify:**
```bash
echo $UV_CACHE_DIR
ls -la $UV_CACHE_DIR
```

---

### Task 5.6: Resolution Strategies
Use different dependency resolution strategies.

```bash
# Lowest compatible versions
uv lock --resolution lowest

# Highest compatible versions (default)
uv lock --resolution highest
```

**Expected Result:**
- Different versions selected
- Useful for testing compatibility range

**Verify:**
```bash
git diff uv.lock
```

---

### Task 5.7: Platform-Specific Dependencies
Add dependencies for specific platforms.

In `pyproject.toml`:
```toml
[project.dependencies]
requests = ">=2.31.0"
pywin32 = { version = ">=305", markers = "platform_system == 'Windows'" }
```

**Expected Result:**
- Dependency only installed on Windows
- Lockfile includes platform markers

**Verify:**
```bash
cat pyproject.toml
uv lock
cat uv.lock | grep pywin32
```

---

### Task 5.8: Create Workspace (Monorepo)
Set up multi-package workspace.

```bash
mkdir -p workspace/{pkg-a,pkg-b}

# Root pyproject.toml
cat > workspace/pyproject.toml << 'EOF'
[tool.uv.workspace]
members = ["pkg-a", "pkg-b"]
EOF

# Create package A
cd workspace/pkg-a
uv init
cd ../..

# Create package B
cd workspace/pkg-b
uv init
cd ../..
```

**Expected Result:**
- Workspace with multiple packages
- Shared lockfile at root
- Packages can depend on each other

**Verify:**
```bash
tree workspace/
cat workspace/pyproject.toml
```

---

### Task 5.9: Frozen Lockfile Install
Install without updating lockfile (CI/CD mode).

```bash
uv sync --frozen
```

**Expected Result:**
- Installs exactly from lockfile
- Fails if lockfile is outdated
- Ensures reproducibility

**Verify:**
```bash
# Try changing pyproject.toml then:
uv sync --frozen  # Should fail/warn
```

---

### Task 5.10: No Dev Dependencies
Install only production dependencies.

```bash
uv sync --no-dev
```

**Expected Result:**
- Dev dependencies not installed
- Smaller environment
- Good for production builds

**Verify:**
```bash
uv pip list | grep pytest  # Should not appear
```

---

## Challenge Task 🏆

Build a complete monorepo project:

**Structure:**
```
my-workspace/
├── pyproject.toml (workspace root)
├── packages/
│   ├── core/
│   │   ├── pyproject.toml
│   │   └── src/core/utils.py
│   ├── api/
│   │   ├── pyproject.toml
│   │   └── src/api/main.py (depends on core)
│   └── cli/
│       ├── pyproject.toml
│       └── src/cli/commands.py (depends on core)
```

**Requirements:**
1. **Core** package with utility functions
2. **API** package (FastAPI) that uses core
3. **CLI** package (Click) that uses core
4. **Workspace** configuration linking all three
5. **Single lockfile** for entire workspace
6. **Scripts** defined for running API and CLI

**Success Criteria:**
- ✅ All packages in workspace
- ✅ `api` and `cli` depend on `core`
- ✅ Single `uv sync` installs everything
- ✅ Changes to `core` reflect in `api` and `cli`
- ✅ Can run both API and CLI from workspace root

---

## Key Takeaways

✅ `uv pip compile` - Resolve and lock dependencies
✅ `uv lock --upgrade` - Update all dependencies
✅ `uv lock --upgrade-package` - Update specific package
✅ `uv sync --frozen` - Install from lockfile strictly
✅ `uv sync --no-dev` - Production-only install
✅ Workspaces - Monorepo support
✅ Scripts - Custom commands in pyproject.toml
✅ Resolution strategies - Control version selection

---

## Next Steps
Move to Exercise 6: Real-World Integration
