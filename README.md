# go-mongodbAPI
A Simple RESTful API Using Golang and MongoDB

## Overview
`go-mongodbAPI` is a lightweight and efficient API built with Golang and MongoDB. It demonstrates how to perform CRUD (Create, Read, Update, Delete) operations and manage data through a RESTful interface, making it an ideal foundation for learning or building scalable back-end services.

---

## Features
- 🛠 **CRUD Operations**: Easily manage TODO items with Create, Read, Update, and Delete functionality.
- 💾 **MongoDB Integration**: Seamless interaction with MongoDB for persistent data storage.
- 🛡 **Error Handling**: Basic validation and error responses for robust API behavior.
- 🚀 **Docker Support**: Quick setup and deployment with Docker.
- 🔄 **Versioned Endpoints**: API organized with versioning (`/api/v1`).

---

## Prerequisites
To run this project, you'll need:
- **Go** (version 1.16 or higher)
- **MongoDB** (local or cloud instance)
- **Docker** (optional but recommended)

---

## Installation
1. **Clone the Repository**:
    ```sh
    git clone https://github.com/nathanfabio/go-mongodbAPI.git
    ```
2. **Navigate to the Project Directory**:
    ```sh
    cd go-mongodbAPI
    ```
3. **Install Dependencies**:
    ```sh
    go mod tidy
    ```

---

## Configuration
1. **Create a `.env` File**:
    - In the project root directory, create a `.env` file to store your environment variables.
2. **Add the Following Variables**:
    ```dotenv
    BINARY=your_binary_name
    DB_CONTAINER_NAME=your_db_container_name
    MONGO_DB=your_database_name
    MONGO_DB_USERNAME=your_database_username
    MONGO_DB_PASSWORD=your_database_password
    ```
    - Replace `your_binary_name` with the name of the Go binary (e.g., `go-mongodbAPI`).
    - Replace `your_db_container_name` with the name of your MongoDB container (e.g., `mongodb`).
    - Replace `your_database_name` with your MongoDB database name (e.g., `todo_db`).
    - Replace `your_database_username` and `your_database_password` with your MongoDB credentials.

---

## Running the API

### Start with Docker (Recommended)
1. **Run Docker Containers**:
    ```sh
    make up
    ```
2. **Start the API Server**:
    ```sh
    make start
    ```

### Run Locally
If you prefer not to use Docker:
1. Start MongoDB on your system.
2. Run the API:
    ```sh
    go run main.go
    ```

---

## API Endpoints

| Method | Endpoint                    | Description                     |
|--------|-----------------------------|---------------------------------|
| GET    | `/api/v1/todos`             | Retrieve all TODOs             |
| GET    | `/api/v1/todos/{id}`        | Retrieve a TODO by ID          |
| POST   | `/api/v1/todos/create`      | Create a new TODO              |
| PUT    | `/api/v1/todos/update/{id}` | Update an existing TODO by ID  |
| DELETE | `/api/v1/todos/delete/{id}` | Delete a TODO by ID            |

### Sample TODO Object
```json
{
    "id": "63d91aebc55e5f001ccf4a45",
    "task": "Buy groceries",
    "completed": false,
    "created_at": "2025-01-20T15:00:00Z",
    "update_at": "2025-01-21T12:00:00Z"
}

## Makefile Commands

```makefile
# Start the Docker containers
make up

# Stop and remove the Docker containers
make down

# Start the API server
make start

# Install Go dependencies
make tidy




## LICENSE

This project is licensed under the MIT License.

Feel free to contribute, fork, or share!