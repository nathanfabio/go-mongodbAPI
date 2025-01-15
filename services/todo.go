package services

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Todo struct {
	ID        string `json:"id,omitempty" bson:"_id,omitempty"`
	Task      string `json:"task,omitempty" bson:"task,omitempty"`
	Completed bool   `json:"completed,omitempty" bson:"completed,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdateAt  time.Time `json:"update_at,omitempty" bson:"update_at,omitempty"`
}


var client *mongo.Client

func NewClient(mongo *mongo.Client) Todo {
	client = mongo

	return Todo{}
}

func returnCollection(collection string) *mongo.Collection {
	return client.Database("mongo_api").Collection(collection)
}

func (t *Todo) InsertTodo(input Todo) error {
	collection := returnCollection("todos")

	_, err := collection.InsertOne(context.TODO(), Todo{
		Task:      input.Task,
		Completed: input.Completed,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	})

	if err != nil {
		log.Println("Error on insert todo: ", err)
		return err
	}

	return nil
}