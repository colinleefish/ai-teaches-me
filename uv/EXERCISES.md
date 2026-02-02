# UV Exercises - Quick Reference

## Structure

```
uv/
├── README.md                    # Main learning guide
├── EXERCISES.md                 # This file - quick reference
├── exercises/
│   ├── README.md               # Exercise overview
│   ├── 01-basics.md           # Virtual envs, pip commands
│   ├── 02-project-management.md   # uv init, add, lock, sync
│   ├── 03-python-versions.md  # Python version management
│   ├── 04-tool-management.md  # Global CLI tools
│   ├── 05-advanced-features.md    # Workspaces, scripts
│   └── 06-real-world-integration.md   # CI/CD, Docker
└── demo-project/
    ├── simple-app/             # Basic Python app
    └── todo-api/               # FastAPI REST API
```

## Quick Start

```bash
# Install UV
curl -LsSf https://astral.sh/uv/install.sh | sh

# Start exercises
cd uv/exercises
cat README.md
```

## Exercise Progression

| # | Topic | Time | Key Skills |
|---|-------|------|------------|
| 1 | Basics | 30m | `venv`, `pip install`, `freeze`, `sync` |
| 2 | Projects | 45m | `init`, `add`, `remove`, `lock` |
| 3 | Python Versions | 30m | `python install/pin`, version management |
| 4 | Tools | 30m | `tool install/run`, global CLI tools |
| 5 | Advanced | 60m | workspaces, scripts, compilation |
| 6 | Production | 90m | CI/CD, Docker, deployment |

**Total:** 4-8 hours

## Commands by Exercise

### Exercise 1: Basics
```bash
uv venv                    # Create virtual environment
uv pip install requests    # Install package
uv pip freeze > req.txt    # Export dependencies
uv pip sync req.txt        # Install exact versions
```

### Exercise 2: Projects
```bash
uv init myapp             # Initialize project
uv add fastapi            # Add dependency
uv add --dev pytest       # Add dev dependency
uv lock                   # Generate lockfile
uv sync                   # Install from lockfile
```

### Exercise 3: Python Versions
```bash
uv python list            # List available versions
uv python install 3.12    # Install Python version
uv python pin 3.12        # Pin project version
uv venv --python 3.12     # Create venv with specific Python
```

### Exercise 4: Tools
```bash
uv tool install ruff      # Install CLI tool globally
uv tool run cowsay hello  # Run without installing
uv tool list              # List installed tools
```

### Exercise 5: Advanced
```bash
uv pip compile pyproject.toml     # Compile requirements
uv lock --upgrade                 # Upgrade all dependencies
uv lock --upgrade-package pkg     # Upgrade specific package
uv sync --frozen                  # Install without updating lock
```

### Exercise 6: Production
```bash
uv sync --frozen --no-dev  # Production install
uv sync --no-cache         # Fresh install without cache
```

## Demo Projects

### Simple App
```bash
cd demo-project/simple-app
uv sync
uv run python main.py
uv run hello  # Runs script from pyproject.toml
```

### TODO API
```bash
cd demo-project/todo-api
uv sync
uv run uvicorn src.main:app --reload
# Visit: http://localhost:8000/docs

# Run tests
uv run pytest
```

## Verification Commands

Check your work:

```bash
# Environment check
uv pip list
which python
python --version

# Project check
cat pyproject.toml
cat uv.lock | head -20
ls -la .venv

# Tool check
uv tool list
which ruff
```

## Common Patterns

### New Project Setup
```bash
uv init myproject
cd myproject
uv add fastapi uvicorn
uv add --dev pytest ruff
uv sync
```

### Dependency Update
```bash
uv lock --upgrade
uv sync
git diff uv.lock
```

### Fresh Install
```bash
git clone <repo>
cd <repo>
uv sync
```

### CI/CD
```yaml
- run: curl -LsSf https://astral.sh/uv/install.sh | sh
- run: uv sync --frozen
- run: uv run pytest
```

### Docker
```dockerfile
COPY --from=ghcr.io/astral-sh/uv:latest /uv /bin/uv
RUN uv sync --frozen --no-cache --no-dev
```

## Challenge Tasks Summary

1. **Exercise 1:** Create script to setup + verify identical environments
2. **Exercise 2:** Build complete TODO API with FastAPI
3. **Exercise 3:** Multi-version testing across Python 3.11-3.13
4. **Exercise 4:** Developer toolchain setup script
5. **Exercise 5:** Monorepo with multiple interdependent packages
6. **Exercise 6:** Production-ready app with full CI/CD

## Next Steps

After completing exercises:

1. ✅ Migrate existing project to UV
2. ✅ Set up UV in your CI/CD
3. ✅ Create Docker images with UV
4. ✅ Share UV with your team
5. ✅ Contribute to UV (it's open source!)

## Resources

- **Main Guide:** `README.md`
- **Exercise Details:** `exercises/README.md`
- **Docs:** https://docs.astral.sh/uv/
- **GitHub:** https://github.com/astral-sh/uv

---

**Start learning: [exercises/01-basics.md](exercises/01-basics.md)** 🚀
