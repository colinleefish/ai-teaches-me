# Exercise 1: UV Basics

## Objective
Learn basic UV commands for package installation and virtual environment management.

---

## Tasks

### Task 1.1: Create Virtual Environment
Create a new virtual environment using UV.

```bash
cd uv/exercises
uv venv my-first-env
```

**Expected Result:**
- `.venv` or `my-first-env` directory created
- Contains `bin/`, `lib/`, `include/` directories
- Python interpreter available in `bin/python`

**Verify:**
```bash
ls my-first-env/
source my-first-env/bin/activate  # or: . my-first-env/bin/activate
which python
```

---

### Task 1.2: Install Packages with UV
Install packages using `uv pip`.

```bash
uv pip install requests
uv pip install "pytest>=7.0.0"
uv pip install httpx rich
```

**Expected Result:**
- Packages installed in seconds (much faster than pip)
- Terminal shows resolution and installation progress
- Packages available for import

**Verify:**
```bash
uv pip list
uv pip show requests
python -c "import requests; print(requests.__version__)"
```

---

### Task 1.3: Freeze Dependencies
Export installed packages to requirements file.

```bash
uv pip freeze > requirements.txt
cat requirements.txt
```

**Expected Result:**
- `requirements.txt` created with all installed packages
- Pinned versions (e.g., `requests==2.31.0`)
- Includes transitive dependencies

**Verify:**
```bash
wc -l requirements.txt  # Should show multiple packages
grep requests requirements.txt
```

---

### Task 1.4: Sync Environment
Create clean environment from requirements file.

```bash
deactivate  # Exit current venv
uv venv sync-test-env
source sync-test-env/bin/activate
uv pip sync requirements.txt
```

**Expected Result:**
- New environment has EXACT packages from requirements.txt
- No extra packages
- Matches previous environment exactly

**Verify:**
```bash
uv pip list
# Should match original environment
```

---

### Task 1.5: Uninstall Packages
Remove packages from environment.

```bash
uv pip uninstall rich
uv pip list
```

**Expected Result:**
- Package removed
- Dependencies may remain if used by other packages

**Verify:**
```bash
python -c "import rich"  # Should fail with ImportError
```

---

## Challenge Task 🏆

Create a script that:
1. Creates a fresh venv
2. Installs `click`, `requests`, and `pydantic`
3. Writes a `requirements.txt`
4. Creates ANOTHER venv and syncs it
5. Verifies both environments are identical

**Hints:**
- Use shell script or Python
- Compare `uv pip list` output from both environments
- Time the installations and compare

---

## Key Takeaways

✅ `uv venv` - Create virtual environments
✅ `uv pip install` - Fast package installation
✅ `uv pip freeze` - Export dependencies
✅ `uv pip sync` - Create reproducible environment
✅ UV is 10-100x faster than standard pip

---

## Next Steps
Move to Exercise 2: Project Management
