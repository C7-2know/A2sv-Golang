package controllers

import (
	"net/http"
	"task_manager/data"
	model "task_manager/models"
	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	response := data.Get_tasks()
	c.JSON(http.StatusOK, response)

}

func GetTask(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	response, err := data.Get_task(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func CreateTask(c *gin.Context) {
	var newTask model.Task
	if err:=c.ShouldBindJSON(&newTask);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	data.Create_task(newTask)
	c.JSON(http.StatusCreated, gin.H{"message":"task created"})
}

func UpdateTask(c *gin.Context) {
	id:=c.Param("id")
	if id==""{
		c.JSON(http.StatusBadRequest, gin.H{"error":"id is required"})
		return
	}
	var updated model.Task
	if err:=c.ShouldBindJSON(&updated);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	data.Update_task(id, updated)
	c.JSON(http.StatusOK, gin.H{"message":"task updated"})
}

func DeleteTask(c *gin.Context){
	id:=c.Param("id")
	if id==""{
		c.JSON(http.StatusBadRequest, gin.H{"error":"id is required"})
		return
	}
	_,err:=data.Get_task(id)
	if err!=nil{
		c.JSON(http.StatusNotFound, gin.H{"message":err.Error()})
		return
	}
	data.Delete_task(id)
	c.JSON(http.StatusOK, gin.H{"message":"task deleted"})
}