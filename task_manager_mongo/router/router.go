package router

import (
	"github.com/gin-gonic/gin"
	control "task_manager/controllers")

func Route() {
	router := gin.Default()
	router.GET("/tasks", control.GetTasks)
	router.GET("/tasks/:id", control.GetTask)
	router.POST("/tasks", control.CreateTask)
	router.PUT("/tasks/:id", control.UpdateTask)
	router.DELETE("/tasks/:id", control.DeleteTask)

	router.Run()
}