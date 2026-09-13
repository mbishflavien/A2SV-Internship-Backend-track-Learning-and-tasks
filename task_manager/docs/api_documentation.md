# Task Management REST API Documentation

Base URL: `http://localhost:8080`

## Endpoints Summary

| Method | Endpoint | Description | Status Code |
| :--- | :--- | :--- | :--- |
| **GET** | `/tasks` | List all tasks | `200 OK` |
| **GET** | `/tasks/:id` | Get single task details | `200 OK` / `404 Not Found` |
| **POST** | `/tasks` | Create a new task | `201 Created` / `400 Bad Request` |
| **PUT** | `/tasks/:id` | Update an existing task | `200 OK` / `400 Bad Request` / `404 Not Found` |
| **DELETE**| `/tasks/:id` | Delete a task by ID | `200 OK` / `404 Not Found` |

---

## Endpoint Details

### 1. Create Task
- **URL**: `/tasks`
- **Method**: `POST`
- **Body**:
```json
{
  "title": "Complete Backend Assignment",
  "description": "Implement Gin REST API task manager",
  "due_date": "2026-09-20T15:04:05Z",
  "status": "Pending"
}