package usecases

import (
	"errors"
	domain "task_manager/Domain"
)


type TaskUsecase interface{
	CreateTask(task *domain.Task) error
	GetAllTask() ([]domain.Task, error)
	GetByIDTask(id int) (domain.Task, error)
}


type taskUsecase struct {
	Task_repo domain.TaskRepository	
}

func NewTaskUsecase(task_repo domain.TaskRepository) domain.TaskUsecase {
	return &taskUsecase{Task_repo: task_repo,}
}

func (tu *taskUsecase) CreateTask(task domain.Task)error{
	err:=tu.Task_repo.CreateTask(task)
	if err!=nil{
		return errors.New("could not create task")
	}
	return nil
}
func (tu *taskUsecase) GetTasks()[]domain.Task{
	tasks:=tu.Task_repo.GetTasks()
	return tasks
}
func (tu *taskUsecase) GetTaskByID(id string)(domain.Task,error){
	task,err:= tu.Task_repo.GetTaskByID(id)
	if err!=nil{
		return domain.Task{},errors.New("Task not found")
	}
	return task,nil
}
func (tu *taskUsecase) UpdateTask(id string,task domain.Task)error{
	_,err:=tu.Task_repo.GetTaskByID(id)
	if err!=nil{
		return errors.New("task not found")
	}
	err= tu.Task_repo.UpdateTask(id,task)
	if err!=nil{
		return errors.New("could not update task")
	}
	return nil
}

func (tu *taskUsecase) DeleteTask(id string)error{
	err:= tu.Task_repo.DeleteTask(id)
	if err!=nil{
		return errors.New("could not delete task")
	}
	return nil
}