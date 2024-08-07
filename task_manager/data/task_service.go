package data

import (
	"errors"
	model "task_manager/models"
)

var tasks = []model.Task{}

func Get_tasks() []model.Task {
	return tasks
}

func Get_task(id string) (model.Task, error) {
	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return model.Task{}, errors.New("Task not found")
}

func Create_task(task model.Task) {
	tasks = append(tasks, task)
}

func Update_task(id string,updated model.Task)  {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i] = updated
			return
		}
	}
}

func Delete_task(id string) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return
		}
	}
	
}
