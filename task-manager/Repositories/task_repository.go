package repository

import (
	"context"
	"errors"
	domain "task_manager/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type taskRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewTaskRepository(db *mongo.Database) domain.TaskRepository {
	return &taskRepository{
		db: db, collection: db.Collection("tasks"),
	}
}

func (tr *taskRepository) CreateTask(task domain.Task) error {
	data := bson.D{
		{"id", task.ID},
		{"description", task.Description},
		{"due_date", task.DueDate},
		{"status", task.Status}}
	if task.ID == "" {
		return errors.New("id is required")
	}
	_, err := tr.collection.InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}
	return nil
}
func (tr *taskRepository) GetTaskByID(id string) (domain.Task, error) {
	filter := bson.D{{"id", id}}
	task := tr.collection.FindOne(context.TODO(), filter)
	if task.Err() != nil {
		return domain.Task{}, task.Err()
	}
	var u domain.Task
	task.Decode(&u)
	return u, nil
}

func (tr *taskRepository) GetTasks() []domain.Task {
	var tasks []domain.Task

	cursor, _ := tr.collection.Find(context.TODO(), bson.D{{}})
	for cursor.Next(context.TODO()) {
		var task domain.Task
		cursor.Decode(&task)
		tasks = append(tasks, task)
	}
	return tasks
}

func (tr *taskRepository) UpdateTask(id string, task domain.Task) error {
	filter := bson.D{{"id", id}}
	up_task:= bson.D{
		{"$set", bson.D{
			{"Id", task.ID},
			{"description", task.Description},
			{"due_date", task.DueDate},
			{"status", task.Status},
		}},
	}

	update, err := tr.collection.UpdateOne(context.TODO(), filter, up_task)
	if err != nil {
		return err
	}
	if update.MatchedCount == 0 {
		return errors.New("no task found")
	}
	return nil
}

func (tr *taskRepository) DeleteTask(id string) error {
	filter := bson.D{{"id", id}}
	result, err := tr.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("no task found")
	}
	return nil
}
