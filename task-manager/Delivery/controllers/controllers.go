package controllers

import (
	"net/http"
	domain "task_manager/Domain"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	User_usecase domain.UserUsecase
}

type TaskController struct {
	Task_usecase domain.TaskUsecase
}

func (uc *UserController) SignUp(c *gin.Context) {
	var user domain.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = uc.User_usecase.CreateUser(user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "user created"})
}

func (uc *UserController) LogIn(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := uc.User_usecase.Login(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", "Bearer "+token, 3600, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (uc *UserController) PromoteUser(c *gin.Context) {
	email := c.Param("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
		return
	}
	err := uc.User_usecase.Promote(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user promoted"})
}

// Task controllers

func (tc *TaskController) GetTasks(c *gin.Context) {
	response := tc.Task_usecase.GetTasks()
	c.JSON(http.StatusOK, response)
}

func (tc *TaskController) GetTask(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	response, err := tc.Task_usecase.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var newTask domain.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := tc.Task_usecase.CreateTask(newTask)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{"message": "task created"})
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	var updated domain.Task
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res:=tc.Task_usecase.UpdateTask(id, updated)
	if res != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "task updated"})
}

func (tc TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	err := tc.Task_usecase.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
}
