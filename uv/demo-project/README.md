# UV Demo Project

This folder contains demo projects for practicing UV exercises.

## Structure

```
demo-project/
├── README.md (this file)
├── simple-app/
│   └── (Exercise 1-2: Basic project)
├── todo-api/
│   └── (Exercise 2 Challenge: FastAPI TODO API)
└── workspace-example/
    └── (Exercise 5: Monorepo workspace)
```

## Quick Start

### Simple App (Exercises 1-2)

```bash
cd simple-app
uv sync
uv run python main.py
```

### TODO API (Exercise 2 Challenge)

```bash
cd todo-api
uv sync
uv run uvicorn src.main:app --reload
```

Open http://localhost:8000/docs for API documentation.

### Workspace Example (Exercise 5)

```bash
cd workspace-example
uv sync  # Installs all packages
uv run --directory packages/api uvicorn src.main:app
```

## Notes

- Each project has its own `pyproject.toml`
- Practice creating/modifying these projects
- Delete and recreate to practice UV commands
- Time yourself - notice UV's speed!
