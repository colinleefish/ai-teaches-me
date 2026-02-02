"""
TODO API - FastAPI demo for UV exercises.
"""

from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from typing import List, Optional

app = FastAPI(title="TODO API", version="1.0.0")

# In-memory storage
todos: List[dict] = []
next_id = 1


class TodoCreate(BaseModel):
    title: str
    description: Optional[str] = None
    completed: bool = False


class Todo(TodoCreate):
    id: int


@app.get("/")
def root():
    """Root endpoint."""
    return {
        "message": "TODO API",
        "docs": "/docs",
        "endpoints": {
            "GET /todos": "List all todos",
            "POST /todos": "Create todo",
            "GET /todos/{id}": "Get todo by ID",
            "PUT /todos/{id}": "Update todo",
            "DELETE /todos/{id}": "Delete todo",
        }
    }


@app.get("/todos", response_model=List[Todo])
def list_todos():
    """List all todos."""
    return todos


@app.post("/todos", response_model=Todo, status_code=201)
def create_todo(todo: TodoCreate):
    """Create a new todo."""
    global next_id
    new_todo = {
        "id": next_id,
        **todo.model_dump()
    }
    todos.append(new_todo)
    next_id += 1
    return new_todo


@app.get("/todos/{todo_id}", response_model=Todo)
def get_todo(todo_id: int):
    """Get a specific todo."""
    for todo in todos:
        if todo["id"] == todo_id:
            return todo
    raise HTTPException(status_code=404, detail="Todo not found")


@app.put("/todos/{todo_id}", response_model=Todo)
def update_todo(todo_id: int, todo_update: TodoCreate):
    """Update a todo."""
    for i, todo in enumerate(todos):
        if todo["id"] == todo_id:
            updated = {
                "id": todo_id,
                **todo_update.model_dump()
            }
            todos[i] = updated
            return updated
    raise HTTPException(status_code=404, detail="Todo not found")


@app.delete("/todos/{todo_id}", status_code=204)
def delete_todo(todo_id: int):
    """Delete a todo."""
    for i, todo in enumerate(todos):
        if todo["id"] == todo_id:
            todos.pop(i)
            return
    raise HTTPException(status_code=404, detail="Todo not found")


@app.get("/health")
def health():
    """Health check."""
    return {"status": "healthy", "todos_count": len(todos)}
