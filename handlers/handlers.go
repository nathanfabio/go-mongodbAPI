package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// healthCheck is a simple handler to check if the server is up and running
func healthCheck(w http.ResponseWriter, r *http.Request) {
	res := Response{
		Message: "Server is up and running",
		Code:    200,
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