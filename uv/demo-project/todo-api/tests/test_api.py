"""
Tests for TODO API
"""

from fastapi.testclient import TestClient
from src.main import app

client = TestClient(app)


def test_root():
    """Test root endpoint."""
    response = client.get("/")
    assert response.status_code == 200
    assert "message" in response.json()


def test_health():
    """Test health check."""
    response = client.get("/health")
    assert response.status_code == 200
    assert response.json()["status"] == "healthy"


def test_create_todo():
    """Test creating a todo."""
    response = client.post(
        "/todos",
        json={"title": "Test Todo", "description": "Test Description"}
    )
    assert response.status_code == 201
    data = response.json()
    assert data["title"] == "Test Todo"
    assert "id" in data


def test_list_todos():
    """Test listing todos."""
    response = client.get("/todos")
    assert response.status_code == 200
    assert isinstance(response.json(), list)


def test_get_todo():
    """Test getting a specific todo."""
    # Create a todo first
    create_response = client.post(
        "/todos",
        json={"title": "Get Test"}
    )
    todo_id = create_response.json()["id"]
    
    # Get it
    response = client.get(f"/todos/{todo_id}")
    assert response.status_code == 200
    assert response.json()["title"] == "Get Test"


def test_update_todo():
    """Test updating a todo."""
    # Create a todo first
    create_response = client.post(
        "/todos",
        json={"title": "Update Test", "completed": False}
    )
    todo_id = create_response.json()["id"]
    
    # Update it
    response = client.put(
        f"/todos/{todo_id}",
        json={"title": "Updated", "completed": True}
    )
    assert response.status_code == 200
    data = response.json()
    assert data["title"] == "Updated"
    assert data["completed"] is True


def test_delete_todo():
    """Test deleting a todo."""
    # Create a todo first
    create_response = client.post(
        "/todos",
        json={"title": "Delete Test"}
    )
    todo_id = create_response.json()["id"]
    
    # Delete it
    response = client.delete(f"/todos/{todo_id}")
    assert response.status_code == 204
    
    # Verify it's gone
    response = client.get(f"/todos/{todo_id}")
    assert response.status_code == 404
