# UV Exercises

Comprehensive hands-on exercises to learn UV, the modern Python package manager.

## Overview

These exercises will take you from UV basics to production-ready workflows.

### Exercise Structure

Each exercise includes:
- 🎯 **Clear objectives**
- 📝 **Step-by-step tasks**
- ✅ **Expected results**
- 🔍 **Verification commands**
- 🏆 **Challenge tasks**

---

## Exercise List

### 1. [Basics](./01-basics.md)
**Time:** 30 minutes  
**Topics:** Virtual environments, package installation, basic commands  
**Skills:** `uv venv`, `uv pip install`, `uv pip freeze`, `uv pip sync`

### 2. [Project Management](./02-project-management.md)
**Time:** 45 minutes  
**Topics:** Modern project structure, dependencies, lockfiles  
**Skills:** `uv init`, `uv add`, `uv remove`, `uv lock`, `uv sync`

### 3. [Python Version Management](./03-python-versions.md)
**Time:** 30 minutes  
**Topics:** Installing and managing Python versions  
**Skills:** `uv python install`, `uv python pin`, `uv venv --python`

### 4. [Tool Management](./04-tool-management.md)
**Time:** 30 minutes  
**Topics:** Global CLI tools, replacng pipx  
**Skills:** `uv tool install`, `uv tool run`, `uv tool list`

### 5. [Advanced Features](./05-advanced-features.md)
**Time:** 60 minutes  
**Topics:** Workspaces, scripts, resolution strategies  
**Skills:** `uv pip compile`, `uv lock --upgrade`, workspaces, `pyproject.toml` scripts

### 6. [Real-World Integration](./06-real-world-integration.md)
**Time:** 90 minutes  
**Topics:** CI/CD, Docker, production deployment  
**Skills:** GitHub Actions, Docker multi-stage builds, deployment automation

---

## Getting Started

### Prerequisites

1. **Install UV:**
   ```bash
   curl -LsSf https://astral.sh/uv/install.sh | sh
   ```

2. **Verify installation:**
   ```bash
   uv --version
   ```

3. **Navigate to exercises:**
   ```bash
   cd uv/exercises
   ```

### Recommended Order

Work through exercises **in order** (1 → 6):
- Each builds on previous concepts
- Difficulty increases gradually
- Later exercises reference earlier setups

### Time Commitment

- **Minimum:** 4 hours (tasks only)
- **Recommended:** 6-8 hours (with challenges)
- **Complete:** 10-12 hours (with all bonus tasks)

---

## Demo Projects

Practice with pre-built demo projects in `../demo-project/`:

### Simple App
Basic Python application for exercises 1-2.

```bash
cd ../demo-project/simple-app
uv sync
uv run python main.py
```

### TODO API
FastAPI REST API for exercises 2-6.

```bash
cd ../demo-project/todo-api
uv sync
uv run uvicorn src.main:app --reload
```

Visit: http://localhost:8000/docs

---

## Tips for Success

### 1. **Experiment Freely**
- Delete and recreate projects
- Try variations of commands
- Break things and fix them

### 2. **Time Yourself**
- Compare UV vs pip speed
- Notice faster workflows
- Track improvement

### 3. **Use Documentation**
- Refer to main README.md
- Check official docs: https://docs.astral.sh/uv/
- Search for specific features

### 4. **Take Notes**
- Document what you learn
- Save useful commands
- Note gotchas and tips

### 5. **Complete Challenges**
- Challenge tasks reinforce learning
- Build real-world skills
- Create portfolio projects

---

## Progress Tracking

Mark your progress:

- [ ] Exercise 1: Basics
- [ ] Exercise 2: Project Management
- [ ] Exercise 3: Python Version Management
- [ ] Exercise 4: Tool Management
- [ ] Exercise 5: Advanced Features
- [ ] Exercise 6: Real-World Integration

---

## Common Issues

### UV Not Found
```bash
# Add to PATH (should be in shell config)
export PATH="$HOME/.local/bin:$PATH"
```

### Cache Issues
```bash
# Clear UV cache
rm -rf ~/.cache/uv
# Or use environment variable
export UV_NO_CACHE=1
```

### Permission Errors
```bash
# UV installs user-level, no sudo needed
# If you see permission errors, don't use sudo
```

---

## What You'll Learn

By completing these exercises, you'll be able to:

✅ Replace pip, pipx, pyenv, and poetry with UV  
✅ Manage Python projects with modern best practices  
✅ Create reproducible environments with lockfiles  
✅ Install and manage multiple Python versions  
✅ Set up monorepo workspaces  
✅ Integrate UV in CI/CD pipelines  
✅ Build optimized Docker images  
✅ Deploy production applications with UV  
✅ Work 10-100x faster than traditional tools

---

## Additional Resources

- **Main README:** `../README.md`
- **Official Docs:** https://docs.astral.sh/uv/
- **GitHub:** https://github.com/astral-sh/uv
- **Astral Blog:** https://astral.sh/blog

---

## Getting Help

If you get stuck:

1. Check the verification commands
2. Review the main README
3. Read error messages carefully
4. Check official documentation
5. Experiment with variations

Remember: The best way to learn is by doing! 🚀

---

**Ready to start? Begin with [Exercise 1: Basics](./01-basics.md)**
