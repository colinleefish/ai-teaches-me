# Exercise 4: Tool Management

## Objective
Learn to use UV to install and run CLI tools globally (replaces pipx).

---

## Tasks

### Task 4.1: Install Global Tool
Install a Python CLI tool globally.

```bash
uv tool install ruff
uv tool install black
uv tool install httpie
```

**Expected Result:**
- Tools installed in isolated environments
- Available globally as commands
- Don't pollute system Python

**Verify:**
```bash
which ruff
ruff --version
which black
black --version
```

---

### Task 4.2: List Installed Tools
See all globally installed tools.

```bash
uv tool list
```

**Expected Result:**
- Lists all tools installed via `uv tool install`
- Shows versions
- Shows install paths

**Verify:**
```bash
uv tool list | grep ruff
```

---

### Task 4.3: Run Tool Without Installing
Execute a tool without permanent installation.

```bash
uv tool run cowsay "Hello from UV!"
uv tool run rich --version
```

**Expected Result:**
- Tool downloads temporarily
- Executes command
- No permanent installation
- Cached for subsequent runs

**Verify:**
```bash
# Run again - should be instant (cached)
uv tool run cowsay "Testing cache"
```

---

### Task 4.4: Upgrade Tool
Update an installed tool to latest version.

```bash
uv tool upgrade ruff
```

**Expected Result:**
- Tool upgraded to latest version
- Shows old → new version

**Verify:**
```bash
ruff --version
uv tool list | grep ruff
```

---

### Task 4.5: Uninstall Tool
Remove globally installed tool.

```bash
uv tool uninstall black
```

**Expected Result:**
- Tool removed
- Command no longer available

**Verify:**
```bash
which black  # Should not be found
uv tool list  # Should not show black
```

---

### Task 4.6: Install Tool with Specific Version
Install tool at specific version.

```bash
uv tool install "ruff==0.1.0"
```

**Expected Result:**
- Specific version installed
- Pinned to that version

**Verify:**
```bash
ruff --version  # Should show 0.1.0
```

---

### Task 4.7: Run Tool from URL
Execute a Python script directly from URL.

```bash
uv tool run https://gist.githubusercontent.com/example/script.py
```

**Expected Result:**
- Script downloaded and executed
- Dependencies installed temporarily

**Note:** Find a real example script or create your own for this task.

---

## Challenge Task 🏆

Create a developer toolchain setup script:

1. **Install** these tools globally:
   - `ruff` (linter)
   - `black` (formatter)
   - `mypy` (type checker)
   - `pytest` (test runner)
   - `httpie` (HTTP client)

2. **Create** a verification script that:
   - Lists all installed tools
   - Shows version of each
   - Tests each tool with a simple command

3. **Bonus**: Create uninstall script that removes all tools

**Success Criteria:**
- ✅ All tools install successfully
- ✅ All tools are globally available
- ✅ Verification script passes
- ✅ Tools work in any directory

**Example verification:**
```bash
ruff check .
black --check .
mypy --version
pytest --version
http --version
```

---

## Key Takeaways

✅ `uv tool install` - Install CLI tools globally
✅ `uv tool run` - Run tool without installing
✅ `uv tool list` - List installed tools
✅ `uv tool upgrade` - Update tool
✅ `uv tool uninstall` - Remove tool
✅ UV replaces pipx for tool management
✅ Each tool has isolated environment

---

## Next Steps
Move to Exercise 5: Advanced Features
