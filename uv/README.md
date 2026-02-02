# UV - Modern Python Package Management

A comprehensive guide to UV, the extremely fast Python package installer and resolver written in Rust by Astral (creators of Ruff).

## Table of Contents
1. [What is UV?](#what-is-uv)
2. [Installation](#installation)
3. [Core Concepts](#core-concepts)
4. [Basic Usage](#basic-usage)
5. [Project Management](#project-management)
6. [Advanced Features](#advanced-features)
7. [Migration Guide](#migration-guide)
8. [Best Practices](#best-practices)
9. [🎯 Hands-On Exercises](#hands-on-exercises)

---

## What is UV?

UV is a **10-100x faster** alternative to pip and pip-tools, designed to:
- Replace pip, pip-tools, pipx, poetry, pyenv, virtualenv
- Provide a unified interface for Python package and project management
- Offer resolution speeds comparable to modern JavaScript/Rust package managers

### Key Features
- ⚡ **Blazingly fast** - written in Rust, 10-100x faster than pip
- 🔒 **Lockfiles** - reproducible dependencies via `uv.lock`
- 🐍 **Python version management** - download and manage Python installations
- 📦 **All-in-one** - replaces multiple tools (pip, virtualenv, pyenv, poetry)
- 🔄 **Drop-in replacement** - mostly compatible with pip commands

---

## Installation

### macOS/Linux
```bash
curl -LsSf https://astral.sh/uv/install.sh | sh
```

### Windows
```powershell
powershell -c "irm https://astral.sh/uv/install.ps1 | iex"
```

### Via Homebrew
```bash
brew install uv
```

### Via pip (not recommended)
```bash
pip install uv
```

### Verify installation
```bash
uv --version
```

---

## Core Concepts

### 1. Virtual Environments
UV automatically manages virtual environments:
```bash
# Create venv
uv venv

# Create with specific Python version
uv venv --python 3.12
```

### 2. Dependencies
UV uses `pyproject.toml` for dependency specification and generates `uv.lock` for reproducibility.

### 3. Python Version Management
UV can install and manage Python versions:
```bash
# List available Python versions
uv python list

# Install specific version
uv python install 3.12

# Use specific version
uv python pin 3.12
```

---

## Basic Usage

### Package Installation

```bash
# Install packages (like pip install)
uv pip install requests

# Install from requirements.txt
uv pip install -r requirements.txt

# Install with specific version
uv pip install "django>=4.0,<5.0"

# Install in development mode
uv pip install -e .
```

### Package Management

```bash
# List installed packages
uv pip list

# Freeze dependencies
uv pip freeze > requirements.txt

# Uninstall package
uv pip uninstall requests

# Sync environment to requirements
uv pip sync requirements.txt
```

### Running Commands

```bash
# Run script with uv
uv run python script.py

# Run with specific Python version
uv run --python 3.12 python script.py
```

---

## Project Management

### Initialize New Project

```bash
# Create new project
uv init my-project
cd my-project

# Structure created:
# my-project/
#   ├── pyproject.toml
#   ├── README.md
#   └── src/
```

### pyproject.toml Example

```toml
[project]
name = "my-project"
version = "0.1.0"
description = "My awesome project"
requires-python = ">=3.12"
dependencies = [
    "requests>=2.31.0",
    "pydantic>=2.0.0",
]

[project.optional-dependencies]
dev = [
    "pytest>=7.0.0",
    "ruff>=0.1.0",
]

[tool.uv]
dev-dependencies = [
    "pytest>=7.0.0",
]
```

### Add Dependencies

```bash
# Add dependency
uv add requests

# Add dev dependency
uv add --dev pytest

# Add with version constraint
uv add "fastapi>=0.104.0"
```

### Remove Dependencies

```bash
uv remove requests
```

### Lock Dependencies

```bash
# Generate/update uv.lock
uv lock

# Install from lockfile
uv sync
```

---

## Advanced Features

### 1. Dependency Resolution

```bash
# Compile requirements with resolution
uv pip compile pyproject.toml -o requirements.txt

# Upgrade all dependencies
uv lock --upgrade

# Upgrade specific package
uv lock --upgrade-package requests
```

### 2. Tool Installation (like pipx)

```bash
# Install CLI tool globally
uv tool install ruff

# Run tool without installing
uv tool run ruff check .

# List installed tools
uv tool list

# Uninstall tool
uv tool uninstall ruff
```

### 3. Build and Publish

```bash
# Build package
uv build

# Publish to PyPI
uv publish
```

### 4. Scripts

Define scripts in `pyproject.toml`:

```toml
[project.scripts]
my-cli = "my_project.cli:main"
```

Run with:
```bash
uv run my-cli
```

### 5. Workspaces (Monorepo)

```toml
# Root pyproject.toml
[tool.uv.workspace]
members = ["packages/*"]
```

Structure:
```
workspace/
├── pyproject.toml
└── packages/
    ├── package-a/
    │   └── pyproject.toml
    └── package-b/
        └── pyproject.toml
```

---

## Migration Guide

### From pip + requirements.txt

```bash
# Convert requirements.txt to pyproject.toml
uv init

# Add dependencies from requirements.txt
uv add $(cat requirements.txt)

# Or directly sync
uv pip sync requirements.txt
```

### From Poetry

```bash
# Poetry dependencies in pyproject.toml work with UV
uv sync

# Generate lockfile
uv lock
```

### From Pipenv

```bash
# Convert Pipfile to requirements
pipenv requirements > requirements.txt

# Use with UV
uv pip install -r requirements.txt
```

---

## Best Practices

### 1. Use Lockfiles

Always commit `uv.lock` for reproducible builds:
```bash
# Generate lockfile
uv lock

# Install exact versions
uv sync
```

### 2. Pin Python Version

```bash
# Pin in project
uv python pin 3.12

# Creates .python-version file
```

### 3. Separate Dev Dependencies

```toml
[project.optional-dependencies]
dev = [
    "pytest",
    "ruff",
    "mypy",
]
```

Install with:
```bash
uv sync --extra dev
```

### 4. CI/CD Integration

```yaml
# GitHub Actions example
- name: Install uv
  run: curl -LsSf https://astral.sh/uv/install.sh | sh

- name: Install dependencies
  run: uv sync

- name: Run tests
  run: uv run pytest
```

### 5. Docker Integration

```dockerfile
FROM python:3.12-slim

# Install uv
COPY --from=ghcr.io/astral-sh/uv:latest /uv /bin/uv

# Copy project files
COPY pyproject.toml uv.lock ./
RUN uv sync --frozen --no-cache

COPY . .
CMD ["uv", "run", "python", "app.py"]
```

---

## Common Commands Cheat Sheet

| Task | UV Command | Equivalent |
|------|------------|------------|
| Create venv | `uv venv` | `python -m venv` |
| Install package | `uv pip install pkg` | `pip install pkg` |
| Install from file | `uv pip install -r requirements.txt` | `pip install -r requirements.txt` |
| Add dependency | `uv add pkg` | `poetry add pkg` |
| Lock dependencies | `uv lock` | `poetry lock` |
| Sync environment | `uv sync` | `poetry install` |
| Run command | `uv run python script.py` | `poetry run python script.py` |
| Install tool | `uv tool install pkg` | `pipx install pkg` |
| Install Python | `uv python install 3.12` | `pyenv install 3.12` |

---

## Resources

- **Official Docs**: https://docs.astral.sh/uv/
- **GitHub**: https://github.com/astral-sh/uv
- **Announcement**: https://astral.sh/blog/uv
- **Ruff** (same team): https://github.com/astral-sh/ruff

---

## Quick Start Example

```bash
# Install UV
curl -LsSf https://astral.sh/uv/install.sh | sh

# Create new project
uv init my-app
cd my-app

# Add dependencies
uv add fastapi uvicorn

# Create main.py
cat << 'EOF' > src/main.py
from fastapi import FastAPI

app = FastAPI()

@app.get("/")
def read_root():
    return {"Hello": "World"}
EOF

# Run the app
uv run uvicorn src.main:app --reload
```

---

## Why UV?

### Speed Comparison
- **pip**: 10-30 seconds
- **poetry**: 15-45 seconds  
- **UV**: 0.5-2 seconds

### Benefits
✅ Single tool replaces many
✅ Faster development workflow
✅ Better dependency resolution
✅ Modern, actively developed
✅ Compatible with existing Python ecosystem

### When to Use UV
- ✅ New projects
- ✅ Want faster CI/CD
- ✅ Need better dependency management
- ✅ Working with monorepos
- ✅ Want unified Python tooling

### When to Stick with pip/poetry
- ⚠️ Complex custom build systems
- ⚠️ Legacy projects with deep pip integration
- ⚠️ Team not ready to adopt new tools

---

## Hands-On Exercises

Ready to practice? Check out the comprehensive exercise suite:

### 📚 [Start Exercises](./exercises/README.md)

**6 Progressive Exercises covering:**
1. ✅ **Basics** - Virtual environments, package installation (30 min)
2. ✅ **Project Management** - Modern Python projects with lockfiles (45 min)
3. ✅ **Python Versions** - Multi-version management (30 min)
4. ✅ **Tool Management** - Global CLI tools (30 min)
5. ✅ **Advanced Features** - Workspaces, scripts, resolution (60 min)
6. ✅ **Real-World Integration** - CI/CD, Docker, production (90 min)

**Each exercise includes:**
- 🎯 Clear objectives and step-by-step tasks
- ✅ Expected results and verification commands
- 🏆 Challenge tasks for deeper learning
- 📁 Demo projects for hands-on practice

### Quick Start Exercises

```bash
# Navigate to exercises
cd exercises

# Read overview
cat README.md

# Start with Exercise 1
cat 01-basics.md
```

### Demo Projects

Practice with pre-built projects:

```bash
# Simple Python app
cd demo-project/simple-app
uv sync
uv run python main.py

# FastAPI REST API
cd demo-project/todo-api
uv sync
uv run uvicorn src.main:app --reload
```

**[View Exercise Quick Reference →](./EXERCISES.md)**

---

**Happy fast Python development! 🚀**
