package test

import (
	"context"
	domain "task_manager/Domain"
	repository "task_manager/Repositories"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepositoryTestSuite struct {
	suite.Suite
	db *mongo.Database
}

func (suite *TaskRepositoryTestSuite) SetupTest() {
	options := options.Client().ApplyURI("mongodb://localhost:27017")

	client, _ := mongo.Connect(context.Background(), options)
	db := client.Database("test")
	suite.db = db

}

func (suite *TaskRepositoryTestSuite) TestCreateTask() {
	task := domain.Task{
		ID:          "2",
		Description: "test",
		DueDate:     time.Now(),
		Status:      "pending",
	}
	taskRepo := repository.NewTaskRepository(suite.db)
	err := taskRepo.CreateTask(task)
	assert.NoError(suite.T(), err)
}

func (suite *TaskRepositoryTestSuite) TestCreateTask_fail() {
	task := domain.Task{
		Description: "test",
		Status:      "pending",
	}
	taskRepo := repository.NewTaskRepository(suite.db)
	err := taskRepo.CreateTask(task)
	assert.Error(suite.T(), err)
}

func (suite *TaskRepositoryTestSuite) TestGetTasks_Success() {
	taskRepo := repository.NewTaskRepository(suite.db)
	result := taskRepo.GetTasks()
	assert.GreaterOrEqual(suite.T(), len(result), 0)
}

func (suite *TaskRepositoryTestSuite) TestGetTaskByID_Success() {
	taskRepo := repository.NewTaskRepository(suite.db)
	_, err := taskRepo.GetTaskByID("2")
	assert.NoError(suite.T(), err)
}
func (suite *TaskRepositoryTestSuite) TestGetTaskByID_fail() {
	taskRepo := repository.NewTaskRepository(suite.db)
	_, err := taskRepo.GetTaskByID("1")
	assert.Error(suite.T(), err)
}

func (suite *TaskRepositoryTestSuite) TestUpdateTask_Success() {
	task := domain.Task{
		ID:          "2",
		Description: "test",
		DueDate:     time.Now(),
		Status:      "pending",
	}
	taskRepo := repository.NewTaskRepository(suite.db)
	err := taskRepo.UpdateTask("2", task)
	assert.NoError(suite.T(), err)
}

func (suite *TaskRepositoryTestSuite) TestUpdateTask_fail() {
	task := domain.Task{
		ID:          "1",
		Description: "test",
		Status:      "pending",
	}
	taskRepo := repository.NewTaskRepository(suite.db)
	err := taskRepo.UpdateTask("1", task)
	assert.Error(suite.T(), err)
}

func (suite *TaskRepositoryTestSuite) TestZDeleteTask_Success() {
	taskRepo := repository.NewTaskRepository(suite.db)
	err := taskRepo.DeleteTask("2")
	assert.NoError(suite.T(), err)
}

func (suite *TaskRepositoryTestSuite) TestDeleteTask_fail() {
	taskRepo := repository.NewTaskRepository(suite.db)
	err := taskRepo.DeleteTask("1")
	assert.Error(suite.T(), err)
}
func TestTaskRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TaskRepositoryTestSuite))
}

func (suite *TaskRepositoryTestSuite) TearDownTest() {
	suite.db.Drop(context.TODO())
}
