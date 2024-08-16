package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"task_manager/Delivery/controllers"
	domain "task_manager/Domain"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type mockUserUsecase struct {
	mock.Mock
}

func (m *mockUserUsecase) CreateUser(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *mockUserUsecase) GetUserByEmail(email string) (domain.User, error) {
	args := m.Called(email)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *mockUserUsecase) Login(email string, password string) (string, error) {
	args := m.Called(email, password)
	return args.String(0), args.Error(1)
}

func (m *mockUserUsecase) Promote(email string) error {
	args := m.Called(email)
	return args.Error(0)
}

// task usecase mock
type mockTaskUsecase struct {
	mock.Mock
}

func (m *mockTaskUsecase) CreateTask(task domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *mockTaskUsecase) GetTasks() []domain.Task {
	args := m.Called()
	return args.Get(0).([]domain.Task)
}

func (m *mockTaskUsecase) GetTaskByID(id string) (domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *mockTaskUsecase) UpdateTask(id string, task domain.Task) error {
	args := m.Called(id, task)
	return args.Error(0)
}

func (m *mockTaskUsecase) DeleteTask(id string) error {
	args := m.Called(id)
	return args.Error(0)

}

type ControllerTestSuite struct {
	suite.Suite
	mockUserUsecase *mockUserUsecase
	mockTaskUsecase *mockTaskUsecase
	router          *gin.Engine
}

func (suite *ControllerTestSuite) SetupTest() {
	suite.mockUserUsecase = new(mockUserUsecase)
	suite.mockTaskUsecase = new(mockTaskUsecase)
	gin.SetMode(gin.TestMode)
	suite.router = gin.Default()
	user_control := controllers.UserController{User_usecase: suite.mockUserUsecase}
	task_control := controllers.TaskController{Task_usecase: suite.mockTaskUsecase}

	suite.router.POST("/signup", user_control.SignUp)
	suite.router.POST("/login", user_control.LogIn)
	suite.router.PUT("/promote/:email", user_control.PromoteUser)
	suite.router.GET("/tasks", task_control.GetTasks)
	suite.router.GET("/tasks/:id", task_control.GetTask)
	suite.router.POST("/tasks", task_control.CreateTask)
	suite.router.PUT("/tasks/:id", task_control.UpdateTask)
	suite.router.DELETE("/tasks/:id", task_control.DeleteTask)
}

func (suite *ControllerTestSuite) TestGetTasks() {
	suite.mockTaskUsecase.On("GetTasks").Return([]domain.Task{})
	req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	suite.Equal(http.StatusOK, w.Code)
	suite.mockTaskUsecase.AssertExpectations(suite.T())
}
func (suite *ControllerTestSuite) TestSignUp_Success() {
	user := domain.User{Name: "new", Email: "test@example.com", Password: "password",Role:""}
	suite.mockUserUsecase.On("CreateUser", user).Return(nil)
	body := `{"name":"new","email":"test@example.com","password":"password","Role":""}`
	req, _ := http.NewRequest(http.MethodPost, "/signup", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusCreated, w.Code)
}

func (suite *ControllerTestSuite) TestSignUp_Failure() {
	user := domain.User{Name: "new", Email: "test@example.com", Password: "password"}
	suite.mockUserUsecase.On("CreateUser", user).Return(assert.AnError)
	body := `{"name":"new","email":"test@example.com","password":"password"}`
	req, _ := http.NewRequest(http.MethodPost, "/signup", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}

func (suite *ControllerTestSuite) TestLogIn_Success() {
	user := domain.User{Email: "test@example.com", Password: "password"}
	suite.mockUserUsecase.On("Login", user.Email, user.Password).Return("token123", nil)
	body := `{"email":"test@example.com","password":"password"}`
	req, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusOK, w.Code)
}

func (suite *ControllerTestSuite) TestLogIn_Failure() {
	user := domain.User{Email: "test@example.com", Password: "password"}
	suite.mockUserUsecase.On("Login", user.Email, user.Password).Return("", assert.AnError)
	body := `{"email":"test@example.com","password":"password"}`
	req, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}

func (suite *ControllerTestSuite) TestPromoteUser_Success() {
	email := "test@example.com"
	suite.mockUserUsecase.On("Promote", email).Return(nil)
	req, _ := http.NewRequest(http.MethodPut, "/promote/"+email, nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusOK, w.Code)
}


func (suite *ControllerTestSuite) TestCreateTask_Success() {
	task := domain.Task{ID:"01",Title: "new", Description: "test task", Status: "incomplete"}
	suite.mockTaskUsecase.On("CreateTask", task).Return(nil)
	body := `{"id":"01","title":"new","description":"test task","status":"incomplete"}`
	req, _ := http.NewRequest(http.MethodPost, "/tasks", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusCreated, w.Code)
}

func (suite *ControllerTestSuite) TestCreateTask_Failure() {
	task := domain.Task{ID:"01",Title: "new", Description: "test task", Status: "incomplete"}
	suite.mockTaskUsecase.On("CreateTask", task).Return(assert.AnError)
	body := `{"id":"01","title":"new","description":"test task","status":"incomplete"}`
	req, _ := http.NewRequest(http.MethodPost, "/tasks", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}

func (suite *ControllerTestSuite) TestGetTask_Success() {
	task := domain.Task{ID:"01",Title: "new", Description: "test task", Status: "incomplete"}
	suite.mockTaskUsecase.On("GetTaskByID", "01").Return(task, nil)
	req, _ := http.NewRequest(http.MethodGet, "/tasks/01", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusOK, w.Code)
}

func (suite *ControllerTestSuite) TestGetTask_Failure() {
	task := domain.Task{ID:"01",Title: "new", Description: "test task", Status: "incomplete"}
	suite.mockTaskUsecase.On("GetTaskByID", "01").Return(task, assert.AnError)
	req, _ := http.NewRequest(http.MethodGet, "/tasks/01", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusNotFound, w.Code)
}
func (suite *ControllerTestSuite) TestUpdateTask_Success() {
	task := domain.Task{ID:"01",Title: "new", Description: "test task", Status: "incomplete"}
	suite.mockTaskUsecase.On("UpdateTask", "01", task).Return(nil)
	body := `{"id":"01","title":"new","description":"test task","status":"incomplete"}`
	req, _ := http.NewRequest(http.MethodPut, "/tasks/01", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusOK, w.Code)
}

func (suite *ControllerTestSuite) TestUpdateTask_Failure() {
	task := domain.Task{ID:"01",Title: "new", Description: "test task", Status: "incomplete"}
	suite.mockTaskUsecase.On("UpdateTask", "01", task).Return(assert.AnError)
	body := `{"id":"01","title":"new","description":"test task","status":"incomplete"}`
	req, _ := http.NewRequest(http.MethodPut, "/tasks/01", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}

func (suite *ControllerTestSuite) TestDeleteTask_Success() {
	suite.mockTaskUsecase.On("DeleteTask", "01").Return(nil)
	req, _ := http.NewRequest(http.MethodDelete, "/tasks/01", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusOK, w.Code)
}
func (suite *ControllerTestSuite) TestDeleteTask_Failure() {
	suite.mockTaskUsecase.On("DeleteTask", "01").Return(assert.AnError)
	req, _ := http.NewRequest(http.MethodDelete, "/tasks/01", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}


func TestControllerSuite(t *testing.T) {
	suite.Run(t, new(ControllerTestSuite))
}


