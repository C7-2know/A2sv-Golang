package test

import (
	"errors"
	domain "task_manager/Domain"
	usecases "task_manager/Usecases"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) CreateTask(task domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) GetTasks() []domain.Task {
	args := m.Called()
	return args.Get(0).([]domain.Task)
}

func (m *MockTaskRepository) GetTaskByID(id string) (domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *MockTaskRepository) UpdateTask(id string, task domain.Task) error {
	args := m.Called(id, task)
	return args.Error(0)
}

func (m *MockTaskRepository) DeleteTask(id string) error {
	args := m.Called(id)
	return args.Error(0)
}



// test task usecase


type TaskUsecaseSuite struct {
	suite.Suite
	mockRepo *MockTaskRepository
	taskUsecase domain.TaskUsecase
}

func (suite *TaskUsecaseSuite) SetupTest() {
	suite.mockRepo = new(MockTaskRepository)
	suite.taskUsecase = usecases.NewTaskUsecase(suite.mockRepo)
}

func (suite *TaskUsecaseSuite) TestCreateTask_Success() {
	task := domain.Task{ID: "1", Title: "Test Task"}

	suite.mockRepo.On("CreateTask", task).Return(nil)

	err := suite.taskUsecase.CreateTask(task)

	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseSuite) TestCreateTask_Failure() {
	task := domain.Task{ID: "1", Title: "Test Task"}

	suite.mockRepo.On("CreateTask", task).Return(errors.New("could not create task"))

	err := suite.taskUsecase.CreateTask(task)

	suite.EqualError(err, "could not create task")
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseSuite) TestGetTasks_Success() {
	tasks := []domain.Task{{ID: "1", Title: "Test Task"}}

	suite.mockRepo.On("GetTasks").Return(tasks)

	result := suite.taskUsecase.GetTasks()

	suite.Equal(tasks, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseSuite) TestGetTaskByID_Success() {
	task := domain.Task{ID: "1", Title: "Test Task"}

	suite.mockRepo.On("GetTaskByID", "1").Return(task, nil)

	result, err := suite.taskUsecase.GetTaskByID("1")

	suite.NoError(err)
	suite.Equal(task, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseSuite) TestGetTaskByID_Failure() {
	suite.mockRepo.On("GetTaskByID", "2").Return(domain.Task{}, errors.New("Task not found"))

	_, err := suite.taskUsecase.GetTaskByID("2")

	suite.EqualError(err, "Task not found")
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseSuite) TestUpdateTask_Success() {
	task := domain.Task{ID: "1", Title: "Test Task"}

	suite.mockRepo.On("GetTaskByID", "1").Return(task, nil)
	suite.mockRepo.On("UpdateTask", "1", task).Return(nil)

	err := suite.taskUsecase.UpdateTask("1", task)

	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseSuite) TestUpdateTask_Failure() {
	task := domain.Task{ID: "1", Title: "Test Task"}

	suite.mockRepo.On("GetTaskByID", "1").Return(domain.Task{}, errors.New("task not found"))

	err := suite.taskUsecase.UpdateTask("1", task)

	suite.EqualError(err, "task not found")
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseSuite) TestDeleteTask_Success() {
	suite.mockRepo.On("DeleteTask", "1").Return(nil)

	err := suite.taskUsecase.DeleteTask("1")

	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseSuite) TestDeleteTask_Failure() {
	suite.mockRepo.On("DeleteTask", "2").Return(errors.New("could not delete task"))

	err := suite.taskUsecase.DeleteTask("2")

	suite.EqualError(err, "could not delete task")
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseSuite) TearDownTest() {
	suite.mockRepo.AssertExpectations(suite.T())
}
func TestTaskUseCaseTestSuite(t *testing.T) {
    suite.Run(t, new(UserUseCaseTestSuite))
}
