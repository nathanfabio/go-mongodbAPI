package services

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Todo struct {
	ID        string `json:"id,omitempty" bson:"_id,omitempty"`
	Task      string `json:"task,omitempty" bson:"task,omitempty"`
	Completed bool   `json:"completed" bson:"completed"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdateAt  time.Time `json:"update_at,omitempty" bson:"update_at,omitempty"`
}


var client *mongo.Client

// NewClient creates a new client
func NewClient(mongo *mongo.Client) Todo {
	client = mongo

	return Todo{}
}

// returnCollection returns a collection
func returnCollection(collection string) *mongo.Collection {
	return client.Database("mongo_api").Collection(collection)
}

// GetAllTodos returns all todos from the database
func (t *Todo) GetAllTodos() ([]Todo, error) {
	collection := returnCollection("todos")
	var todos []Todo
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var todo Todo
		cursor.Decode(&todo)
		todos = append(todos, todo) 
	}

	return todos, nil
}

// GetTodoByID returns a todo by ID
func (t *Todo) GetTodoByID(id string) (Todo, error) {
	collection := returnCollection("todos")
	var todo Todo

	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Todo{}, err
	}

	err = collection.FindOne(context.Background(), bson.M{"_id": mongoID}).Decode(&todo)
	if err != nil {
		log.Println(err)
		return Todo{}, err
	}

	return todo, nil
}

// UpdateTodo updates a todo in the database
func (t *Todo) UpdateTodo(input Todo) (*mongo.UpdateResult, error) {
	collection := returnCollection("todos")
	mongoID, err := primitive.ObjectIDFromHex(input.ID)
	if err != nil {
		return nil, err
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "task", Value: input.Task},
			{Key: "completed", Value: input.Completed},
			{Key: "update_at", Value: time.Now()},
		}},
	}

	result, err := collection.UpdateOne(context.Background(), bson.M{"_id": mongoID}, update)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}


// DeleteTodo deletes a todo in the database
func (t *Todo) DeleteTodo(id string) error {
	collection := returnCollection("todos")
	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": mongoID})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}


// InsertTodo inserts a todo in the database
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