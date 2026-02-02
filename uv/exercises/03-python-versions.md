# Exercise 3: Python Version Management

## Objective
Learn to manage multiple Python versions with UV (replaces pyenv).

---

## Tasks

### Task 3.1: List Available Python Versions
See what Python versions UV can install.

```bash
uv python list
```

**Expected Result:**
- List of available Python versions
- Shows installed versions (marked)
- Includes CPython, PyPy variants

**Verify:**
```bash
uv python list | grep "3.12"
uv python list | grep "3.11"
```

---

### Task 3.2: Install Specific Python Version
Install a Python version using UV.

```bash
uv python install 3.12
uv python install 3.11
```

**Expected Result:**
- Python downloaded and installed by UV
- Available for UV to use
- Fast download and installation

**Verify:**
```bash
uv python list  # Should show installed versions marked
```

---

### Task 3.3: Create Venv with Specific Python
Create virtual environment with specific Python version.

```bash
uv venv --python 3.12 venv-py312
uv venv --python 3.11 venv-py311
```

**Expected Result:**
- Two separate environments created
- Each with different Python version

**Verify:**
```bash
venv-py312/bin/python --version  # Should show 3.12.x
venv-py311/bin/python --version  # Should show 3.11.x
```

---

### Task 3.4: Pin Python Version in Project
Set required Python version for a project.

```bash
cd uv/demo-project/myapp  # Or any project
uv python pin 3.12
```

**Expected Result:**
- `.python-version` file created
- Contains `3.12`
- UV will use this version automatically

**Verify:**
```bash
cat .python-version
uv venv  # Should automatically use Python 3.12
```

---

### Task 3.5: Find Python Installations
Locate Python installations on your system.

```bash
uv python find 3.12
uv python find 3.11
```

**Expected Result:**
- Prints path to Python executable
- Shows which installation will be used

**Verify:**
```bash
uv python find 3.12 | xargs -I {} {} --version
```

---

### Task 3.6: Use Python for One-Off Commands
Run command with specific Python version.

```bash
uv run --python 3.12 python --version
uv run --python 3.11 python --version
```

**Expected Result:**
- Command runs with specified Python version
- No need to activate environment

**Verify:**
```bash
# Should show different versions
uv run --python 3.12 python -c "import sys; print(sys.version)"
uv run --python 3.11 python -c "import sys; print(sys.version)"
```

---

## Challenge Task 🏆

Create a multi-version testing setup:

1. **Create** a simple Python package with function:
   ```python
   def get_python_version():
       import sys
       return f"{sys.version_info.major}.{sys.version_info.minor}"
   ```

2. **Test** it works on Python 3.11, 3.12, and 3.13

3. **Create** a shell script that:
   - Installs all three Python versions
   - Creates separate venvs for each
   - Runs the function in each environment
   - Prints results

**Success Criteria:**
- ✅ Script runs successfully
- ✅ Shows three different Python versions
- ✅ No manual environment activation needed

**Bonus:**
- Time how long UV takes vs traditional pyenv

---

## Key Takeaways

✅ `uv python list` - List available Python versions
✅ `uv python install` - Install Python versions
✅ `uv python pin` - Set project Python version
✅ `uv python find` - Locate Python installations
✅ `uv venv --python X.Y` - Create venv with specific Python
✅ `uv run --python X.Y` - Run with specific Python
✅ UV replaces pyenv for version management

---

## Next Steps
Move to Exercise 4: Tool Management
