# go-mongodbAPI
A simple API with Golang and MongoDB


## Overview
This project is a simple API built with Golang and MongoDB. It demonstrates how to create a RESTful API using the Go programming language and interact with a MongoDB database.

## Features
- Create, read, update, and delete (CRUD) operations for managing data.
- Connection to MongoDB for data storage.
- Basic error handling and validation.

## Prerequisites
- Go 1.16 or higher
- MongoDB

## Installation
1. Clone the repository:
    ```sh
    git clone https://github.com/nathanfabio/go-mongodbAPI.git
    ```
2. Navigate to the project directory:
    ```sh
    cd go-mongodbAPI
    ```
3. Install the dependencies:
    ```sh
    go mod tidy
    ```

## Configuration
Update the .env file with your MongoDB connection details.

## Running the API
Start the docker containers:
```sh
make up
```

Start the API server:
```sh
make start
```

## API Endpoints
- `GET /api/v1/todos` - Retrieve all TODOS
- `GET /api/v1/todos/{id}` - Retrieve a specific TODO by ID
- `POST /api/v1/create` - Create a new TODO
- `PUT /api/v1/update/{id}` - Update an existing TODO by ID
- `DELETE /api/v1/delete/{id}` - Delete an TODO by ID


## License
This project is licensed under the MIT License.
