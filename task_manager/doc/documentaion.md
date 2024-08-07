Task Manager API Documentation
Overview
The Task Manager API provides endpoints for managing tasks in a task management system. I

- Base URL --- /tasks
Endpoints
1. Get All Tasks
Endpoint: GET /tasks
Description: Retrieves a list of all tasks.

Response: 
    Status Code: 200 OK
    Content-Type: application/json
    Body:   json
        [
        {
            "id": "1",
            "title": "Task 1",
            "description": "Description of Task 1",
            "status": "completed"
        },
        ...
        ]

2. Get Task By ID
Endpoint: GET /tasks/:id
Description: Retrieves a single task by its ID.
Parameters:
    Path Parameter: id (string) - The unique identifier of the task.
Response: 
    Content-Type: application/json
    Status Code: 200 OK
        Content-Type: application/json
        Body: json
        {
            "id": "1",
            "title": "Task 1",
            "description": "Description of Task 1",
            "status": "completed"
        }

    Status Code: 400 Bad Request (if id is missing)
        Body:json {"error": "id is required"}
    Status Code: 404 Not Found (if task with id does not exist)
        Body:json{"message": "Task not found"}

3. Create Task
Endpoint: POST /tasks
Description: Creates a new task.
Request Body:
    Content-Type: application/json
    Body: json
        {
        "title": "New Task",
        "description": "Description of the new task",
        "status": "pending"
        }
Response:
    Content-Type: application/json
    Status Code: 201 Created
        Body:json{"message": "task created"}
    Status Code: 400 Bad Request (if request body is invalid)
        Body:json{"error": "Validation error message"}

4. Update Task
Endpoint: PUT /tasks/:id
Description: Updates an existing task by its ID.
Parameters:
Path Parameter: id (string) - The unique identifier of the task.
Request Body:
    Content-Type: application/json
    Body:json
    {
        "title": "Updated Task Title",
        "description": "Updated description",
        "status": "in-progress"
    }
Response:
    Content-Type: application/json
    Status Code: 200 OK
        Body:
        json{"message": "task updated"}
    Status Code: 400 Bad Request (if id is missing or request body is invalid)
        Body:json{"error": "Validation error message"}

5. Delete Task
Endpoint: DELETE /tasks/:id
Description: Deletes a task by its ID.
Parameters:
    Path Parameter: id (string) - The unique identifier of the task.
Response:
    Content-Type: application/json
    Status Code: 200 OK
        Body:json{"message": "task deleted"}
    Status Code: 400 Bad Request (if id is missing)
        Body:json{"error": "id is required"}
    Status Code: 404 Not Found (if task with id does not exist)
        Body:json{"message": "Task not found"}
