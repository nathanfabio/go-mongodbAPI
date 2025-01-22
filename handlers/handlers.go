package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nathanfabio/go-mongodbAPI/services"
)

var todo services.Todo

// healthCheck is a simple handler to check if the server is up and running
func healthCheck(w http.ResponseWriter, r *http.Request) {
	res := Response{
		Message: "Server is up and running",
		Code:    http.StatusOK,
	}

	jsonResponse, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// GetTodos is a simple handler to get all todos
func getTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := todo.GetAllTodos()
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}


// GetTodoByID is a simple handler to get a todo by ID
func getTodoByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_, err := todo.GetTodoByID(id)
	if err != nil {
		errResp := Response{
			Message: "Error on get todo",
			Code:    http.StatusInternalServerError,
		}
		w.WriteHeader(errResp.Code)
		json.NewEncoder(w).Encode(errResp)
		return
	}

	res := Response {
		Message: "Todo retrieved successfully",
		Code:    http.StatusOK,
	}

	jsonStr, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonStr)
}


// createTodo is a simple handler to create a todo
func createTodo(w http.ResponseWriter, r *http.Request) {
	// Get the todo from the request body
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Fatal(err)
	}

	err = todo.InsertTodo(todo)
	if err != nil {
		errResp := Response{
			Message: "Error on insert todo",
			Code:    http.StatusInternalServerError,
		}
		json.NewEncoder(w).Encode(errResp)
		return
	}

	res := Response {
		Message: "Todo created successfully",
		Code:    http.StatusCreated,
	}

	jsonStr, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	w.Write(jsonStr)
	
}


// UpdateTodo is a simple handler to update a todo
func updateTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	todo, err := todo.GetTodoByID(id)
	if err != nil {
		log.Println(err)
		return
	}
	
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = todo.UpdateTodo(todo)
	if err != nil {
		errResp := Response{
			Message: "Error on update todo",
			Code:    http.StatusInternalServerError,
		}
		json.NewEncoder(w).Encode(errResp)
		w.WriteHeader(errResp.Code)
		return
	}

	res := Response {
		Message: "Todo updated successfully",
		Code:    http.StatusOK,
	}

	jsonStr, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	w.Write(jsonStr)

}


// DeleteTodo is a simple handler to delete a todo
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := todo.DeleteTodo(id)
	if err != nil {
		errResp := Response{
			Message: "Error on delete todo",
			Code:    http.StatusInternalServerError,
		}
		json.NewEncoder(w).Encode(errResp)
		w.WriteHeader(errResp.Code)
		return
	}

	res := Response {
		Message: "Todo deleted successfully",
		Code:    http.StatusOK,
	}

	jsonStr, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	w.Write(jsonStr)
}