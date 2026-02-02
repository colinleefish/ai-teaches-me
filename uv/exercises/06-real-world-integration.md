# Exercise 6: Real-World Integration

## Objective
Learn to integrate UV in CI/CD, Docker, and production workflows.

---

## Tasks

### Task 6.1: GitHub Actions CI
Create GitHub Actions workflow with UV.

Create `.github/workflows/test.yml`:

```yaml
name: Test

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Install uv
        run: curl -LsSf https://astral.sh/uv/install.sh | sh
      
      - name: Add uv to PATH
        run: echo "$HOME/.local/bin" >> $GITHUB_PATH
      
      - name: Install dependencies
        run: uv sync
      
      - name: Run tests
        run: uv run pytest
      
      - name: Run linter
        run: uv run ruff check .
```

**Expected Result:**
- UV installs in CI
- Dependencies install quickly
- Tests run successfully
- Much faster than pip-based CI

**Verify:**
- Commit and push
- Check GitHub Actions tab
- Compare timing with pip-based workflows

---

### Task 6.2: Docker Integration
Create optimized Dockerfile with UV.

Create `Dockerfile`:

```dockerfile
FROM python:3.12-slim

# Install uv
COPY --from=ghcr.io/astral-sh/uv:latest /uv /bin/uv

# Set working directory
WORKDIR /app

# Copy dependency files
COPY pyproject.toml uv.lock ./

# Install dependencies (no cache, frozen lockfile)
RUN uv sync --frozen --no-cache --no-dev

# Copy application code
COPY . .

# Run application
CMD ["uv", "run", "python", "-m", "myapp"]
```

**Expected Result:**
- Small image size
- Fast builds (UV speed)
- Reproducible (frozen lockfile)
- No unnecessary dev dependencies

**Verify:**
```bash
docker build -t myapp .
docker run myapp
docker images | grep myapp  # Check size
```

---

### Task 6.3: Docker Multi-Stage Build
Optimize Docker image with multi-stage build.

```dockerfile
# Stage 1: Build
FROM python:3.12-slim AS builder

COPY --from=ghcr.io/astral-sh/uv:latest /uv /bin/uv

WORKDIR /app

COPY pyproject.toml uv.lock ./
RUN uv sync --frozen --no-dev --no-cache

# Stage 2: Runtime
FROM python:3.12-slim

WORKDIR /app

# Copy only installed packages from builder
COPY --from=builder /app/.venv /app/.venv
COPY . .

ENV PATH="/app/.venv/bin:$PATH"

CMD ["python", "-m", "myapp"]
```

**Expected Result:**
- Even smaller final image
- No build tools in runtime image
- Faster deployments

**Verify:**
```bash
docker build -t myapp-optimized .
docker images | grep myapp
# Compare sizes
```

---

### Task 6.4: Pre-commit Hooks
Set up pre-commit with UV-installed tools.

Create `.pre-commit-config.yaml`:

```yaml
repos:
  - repo: local
    hooks:
      - id: ruff
        name: ruff
        entry: uv run ruff check --fix
        language: system
        types: [python]
      
      - id: black
        name: black
        entry: uv run black
        language: system
        types: [python]
      
      - id: mypy
        name: mypy
        entry: uv run mypy
        language: system
        types: [python]
```

**Expected Result:**
- Code checked before commit
- Uses UV-managed tools
- Fast execution

**Verify:**
```bash
uv tool install pre-commit
pre-commit install
# Make a change and commit
git add .
git commit -m "test"
```

---

### Task 6.5: Makefile Integration
Create Makefile for common UV tasks.

Create `Makefile`:

```makefile
.PHONY: install update test lint format clean

install:
	uv sync

update:
	uv lock --upgrade
	uv sync

test:
	uv run pytest

lint:
	uv run ruff check .

format:
	uv run ruff format .

clean:
	rm -rf .venv uv.lock
	find . -type d -name __pycache__ -exec rm -rf {} +

dev:
	uv run uvicorn myapp.main:app --reload
```

**Expected Result:**
- Simple commands: `make install`, `make test`
- Consistent across team
- Easy onboarding

**Verify:**
```bash
make install
make test
make lint
```

---

### Task 6.6: VS Code Integration
Configure VS Code to use UV environment.

Create `.vscode/settings.json`:

```json
{
  "python.defaultInterpreterPath": "${workspaceFolder}/.venv/bin/python",
  "python.terminal.activateEnvironment": true,
  "python.testing.pytestEnabled": true,
  "python.testing.pytestArgs": ["."],
  "python.linting.enabled": true,
  "python.formatting.provider": "black",
  "[python]": {
    "editor.defaultFormatter": "ms-python.black-formatter",
    "editor.formatOnSave": true,
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    }
  }
}
```

**Expected Result:**
- VS Code uses UV's virtual environment
- Linting and formatting work
- Tests discoverable

**Verify:**
- Open project in VS Code
- Check Python interpreter (bottom right)
- Run tests from Test Explorer

---

### Task 6.7: Production Deployment Script
Create deployment script with UV.

Create `deploy.sh`:

```bash
#!/bin/bash
set -e

echo "🚀 Deploying application..."

# Pull latest code
git pull origin main

# Install dependencies
echo "📦 Installing dependencies..."
uv sync --frozen --no-dev

# Run migrations
echo "🗄️  Running migrations..."
uv run python manage.py migrate

# Collect static files
echo "📁 Collecting static files..."
uv run python manage.py collectstatic --noinput

# Restart service
echo "♻️  Restarting service..."
sudo systemctl restart myapp

echo "✅ Deployment complete!"
```

**Expected Result:**
- Automated deployment
- Fast dependency installation
- Safe with frozen lockfile

**Verify:**
```bash
chmod +x deploy.sh
./deploy.sh  # (in staging environment)
```

---

### Task 6.8: Performance Benchmarking
Compare UV vs pip speed.

Create `benchmark.sh`:

```bash
#!/bin/bash

PROJECT_DIR="/tmp/benchmark"

echo "🏎️  Benchmarking UV vs pip..."

# Benchmark UV
echo "\n=== UV ==="
rm -rf $PROJECT_DIR
mkdir -p $PROJECT_DIR
cd $PROJECT_DIR

time (
  uv venv
  uv pip install fastapi uvicorn sqlalchemy pydantic requests httpx pytest black ruff
)

UV_TIME=$?

# Benchmark pip
echo "\n=== pip ==="
rm -rf $PROJECT_DIR
mkdir -p $PROJECT_DIR
cd $PROJECT_DIR

time (
  python -m venv .venv
  source .venv/bin/activate
  pip install fastapi uvicorn sqlalchemy pydantic requests httpx pytest black ruff
)

echo "\n✅ Benchmark complete!"
```

**Expected Result:**
- UV 10-50x faster
- Clear timing difference
- Quantifiable improvement

**Verify:**
```bash
chmod +x benchmark.sh
./benchmark.sh
```

---

## Challenge Task 🏆

Create a **complete production-ready project setup**:

**Components:**
1. ✅ **FastAPI application** with database
2. ✅ **Docker** setup (multi-stage)
3. ✅ **Docker Compose** for local development
4. ✅ **GitHub Actions** CI/CD
5. ✅ **Makefile** for common tasks
6. ✅ **Pre-commit hooks**
7. ✅ **Documentation** with deployment instructions

**Requirements:**
- All dependencies managed with UV
- Frozen lockfile for reproducibility
- Separate dev/prod dependency groups
- Fast CI pipeline (< 2 min)
- Optimized Docker image (< 100MB)
- One-command local setup
- Automated tests in CI

**Success Criteria:**
```bash
# Local development
make install
make dev

# Testing
make test
make lint

# Production
docker build -t myapp .
docker run -p 8000:8000 myapp

# CI passes all checks
git push  # Triggers CI
```

---

## Key Takeaways

✅ UV integrates seamlessly with CI/CD
✅ Docker images build faster with UV
✅ Development workflows simplified
✅ Production deployments more reliable
✅ Team onboarding easier
✅ Significant speed improvements in pipelines

---

## Congratulations! 🎉

You've completed all UV exercises covering:
1. ✅ Basics (venv, pip commands)
2. ✅ Project management (add, lock, sync)
3. ✅ Python version management
4. ✅ Tool management (global CLI tools)
5. ✅ Advanced features (workspaces, scripts)
6. ✅ Real-world integration (CI/CD, Docker)

You're now ready to use UV in production projects!
