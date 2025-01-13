package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/nathanfabio/go-mongodbAPI/db"
	"github.com/nathanfabio/go-mongodbAPI/handlers"
	"github.com/nathanfabio/go-mongodbAPI/services"
)

type Application struct {
	Models services.Models
}

func main() {
	mongoClient, err := db.ConnMongo()
	if err != nil {
		log.Panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	defer func ()  {
		if err = mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	services.NewClient(mongoClient)

	log.Panicln("Running in port 8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CreateRouter()))
}