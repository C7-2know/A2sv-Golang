package data

import (
	"context"
	"fmt"

	model "task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Dbconnect() error {
	clientOption := options.Client().ApplyURI("Mongodb connection")
	var err error
	client, err = mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		return err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}
	return nil
}

func Get_task(id string) (model.Task, error) {
	collection := client.Database("task_manager").Collection("tasks")
	filter := bson.D{{"id", id}}
	result := collection.FindOne(context.TODO(), filter)
	var task model.Task
	err := result.Decode(&task)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func Get_tasks() []model.Task {
	if client == nil {
		fmt.Println("client is nil")
		return []model.Task{}
	}
	collection := client.Database("task_manager").Collection("tasks")
	
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return []model.Task{}
	}
	fmt.Println("cursor", cursor)
	var tasks []model.Task
	for cursor.Next(context.TODO()) {
		var task model.Task
		err := cursor.Decode(&task)
		if err != nil {
			return nil
		}
		tasks = append(tasks, task)
	}
	return tasks
}

func Create_task(task model.Task) error {
	collection := client.Database("task_manager").Collection("tasks")
	_, err := collection.InsertOne(context.TODO(), task)
	if err != nil {
		return err
	}
	return nil
}

func Update_task(id string, updated model.Task) error {
	collection := client.Database("task_manager").Collection("tasks")
	filter := bson.D{{"id", id}}
	_, err := collection.UpdateByID(context.TODO(), filter, updated)
	if err != nil {
		return err
	}
	return nil
}

func Delete_task(id string) error {
	collection := client.Database("task_manager").Collection("tasks")
	filter := bson.D{{"id", id}}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}
