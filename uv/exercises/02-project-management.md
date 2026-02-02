# Exercise 2: Project Management

## Objective
Learn to use UV for modern Python project management with `pyproject.toml`.

---

## Tasks

### Task 2.1: Initialize New Project
Create a new Python project with UV.

```bash
cd uv/demo-project
uv init myapp
cd myapp
```

**Expected Result:**
- Directory structure created:
  ```
  myapp/
  ├── pyproject.toml
  ├── README.md
  └── hello.py
  ```
- `pyproject.toml` contains basic project metadata

**Verify:**
```bash
tree myapp/
cat myapp/pyproject.toml
```

---

### Task 2.2: Add Dependencies
Add packages to your project using `uv add`.

```bash
uv add fastapi
uv add uvicorn
uv add "pydantic>=2.0.0"
```

**Expected Result:**
- Dependencies added to `pyproject.toml` under `[project.dependencies]`
- `uv.lock` file created/updated
- Packages installed in project environment

**Verify:**
```bash
cat pyproject.toml  # Check dependencies section
cat uv.lock | head -20
uv pip list
```

---

### Task 2.3: Add Dev Dependencies
Add development-only dependencies.

```bash
uv add --dev pytest
uv add --dev ruff
uv add --dev mypy
```

**Expected Result:**
- Dev dependencies in separate section: `[tool.uv.dev-dependencies]`
- Not included in production builds

**Verify:**
```bash
grep -A 5 "dev-dependencies" pyproject.toml
```

---

### Task 2.4: Remove Dependencies
Remove a dependency from project.

```bash
uv remove uvicorn
```

**Expected Result:**
- Package removed from `pyproject.toml`
- Lockfile updated
- Package uninstalled from environment

**Verify:**
```bash
cat pyproject.toml | grep uvicorn  # Should not exist
```

---

### Task 2.5: Lock Dependencies
Generate/update lockfile explicitly.

```bash
uv lock
```

**Expected Result:**
- `uv.lock` updated with current dependency resolution
- Includes all transitive dependencies with exact versions
- Platform-specific if needed

**Verify:**
```bash
ls -lh uv.lock
cat uv.lock | grep -E "^name|^version" | head -20
```

---

### Task 2.6: Sync Environment
Install exact versions from lockfile.

```bash
# Simulate fresh checkout
rm -rf .venv
uv sync
```

**Expected Result:**
- Virtual environment recreated
- Exact packages from `uv.lock` installed
- Reproducible across machines

**Verify:**
```bash
uv pip list
# Compare with uv.lock contents
```

---

### Task 2.7: Sync with Extras
Install with optional dependency groups.

```bash
uv sync --extra dev
```

**Expected Result:**
- Production dependencies installed
- Development dependencies also installed

**Verify:**
```bash
uv pip list | grep -E "pytest|ruff|mypy"
```

---

## Challenge Task 🏆

Create a complete web API project:

1. **Initialize** project called `todo-api`
2. **Add** these dependencies:
   - `fastapi`
   - `sqlalchemy`
   - `python-dotenv`
3. **Add** these dev dependencies:
   - `pytest`
   - `httpx` (for testing)
   - `ruff`
4. **Create** simple FastAPI app in `src/main.py`
5. **Lock** dependencies
6. **Verify** the project can be installed fresh on "another machine"

**Success Criteria:**
- ✅ Clean `pyproject.toml` with proper dependencies
- ✅ Locked `uv.lock` file
- ✅ API runs with `uv run uvicorn src.main:app`
- ✅ Can delete `.venv` and restore with `uv sync`

---

## Key Takeaways

✅ `uv init` - Initialize new projects
✅ `uv add` - Add dependencies to project
✅ `uv add --dev` - Add development dependencies
✅ `uv remove` - Remove dependencies
✅ `uv lock` - Generate lockfile
✅ `uv sync` - Install from lockfile
✅ `pyproject.toml` - Modern Python project definition

---

## Next Steps
Move to Exercise 3: Python Version Management
