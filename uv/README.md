# uv - Fast Python Package & Project Manager

A minicourse on [uv](https://github.com/astral-sh/uv), the Rust-based Python toolchain from Astral.

## What uv Replaces

| Old Tool  | uv Equivalent    |
| --------- | ---------------- |
| pyenv     | `uv python`      |
| pip       | `uv pip`         |
| venv      | `uv venv`        |
| pip-tools | `uv pip compile` |
| pipx      | `uv tool`        |

## Course Outline

1. **Installation & Setup** - Getting started with uv
2. **Python Version Management** - Installing and switching Python versions
3. **Virtual Environments** - Creating and managing venvs
4. **Dependency Management** - Adding packages, lockfiles, pyproject.toml
5. **Project Workflows** - Real-world usage patterns
6. **Migration Guide** - Moving from pyenv/pip/pipx

## Quick Reference

```bash
# Install Python
uv python install 3.12

# Create venv
uv venv --python 3.12

# Install packages
uv pip install requests

# Run scripts
uv run python script.py

# Install CLI tools globally
uv tool install ruff
```

## Resources

- [Official Docs](https://docs.astral.sh/uv/)
- [GitHub](https://github.com/astral-sh/uv)
