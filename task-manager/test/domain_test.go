package test

import (
	"task_manager/Domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUserModel(t *testing.T) {
	task := domain.User{
		Name:     "Test",
		Email:    "test@gmail.com",
		Password: "123456",
		Role:     "admin",
	}
	assert.Equal(t, "Test", task.Name)
	assert.Equal(t, "test@gmail.com", task.Email)
	assert.Equal(t, "123456", task.Password)
	assert.Equal(t, "admin", task.Role)

}

func TestTaskModel(t *testing.T) {
	date := time.Now()
	task := domain.Task{
		ID:          "1",
		Title:       "Test",
		Description: "Test",
		DueDate:     date,
		Status:      "pending",
	}
	assert.Equal(t, "1", task.ID)
	assert.Equal(t, "Test", task.Title)
	assert.Equal(t, "Test", task.Description)
	assert.Equal(t, date, task.DueDate)
	assert.Equal(t, "pending", task.Status)
}
