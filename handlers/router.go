package handlers

import (

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func CreateRouter() *chi.Mux {
	router := chi.NewRouter() 

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTION"},
		AllowedHeaders: []string{"Content-Type", "Authorization", "Accept", "Origin", "X-Requested-With"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
		MaxAge: 300,
	}))

	router.Route("/api", func(router chi.Router) {
		//version 1
		router.Route("/v1", func(router chi.Router) {
			router.Get("/healthcheck", healthCheck)
			router.Get("/todos/{id}", getTodoByID)
			router.Get("/todos", getTodos)
			router.Post("/todos/create", createTodo)
			router.Put("/todos/update/{id}", updateTodo)
			router.Delete("/todos/delete/{id}", deleteTodo)
		})
	
	})
	
	return router
}