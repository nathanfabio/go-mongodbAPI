package handlers

import (
	"encoding/json"
	"log"
	"net/http"

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

	router.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		res := Response {
			Message: "Server is up and running",
			Code: 200,
		}

		jsonResponse, err := json.Marshal(res)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	})

	return router
}