package services

import (
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