package usecases

import domain "task_manager/Domain"


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
	return tu.Task_repo.CreateTask(task)
}
func (tu *taskUsecase) GetTasks()[]domain.Task{
	return tu.Task_repo.GetTasks()
}
func (tu *taskUsecase) GetTaskByID(id string)(domain.Task,error){
	return tu.Task_repo.GetTaskByID(id)
}
func (tu *taskUsecase) UpdateTask(id string,task domain.Task)error{
	return tu.Task_repo.UpdateTask(id,task)
}

func (tu *taskUsecase) DeleteTask(id string)error{
	return tu.Task_repo.DeleteTask(id)
}