package domain

import (
	"time"
)

type Task struct {
	ID          string    `json:"id" bson:"id" unique:"true"`
	Title       string    `json:"title" bson:"title" required:"true"`
	Description string    `json:"description" bson:"description"`
	DueDate     time.Time `json:"due_date" bson:"due_date"`
	Status      string    `json:"status" bson:"status"`
}

type User struct {
	// ID 	 primitive.ObjectID `json:"id" bson:"_id" unique:"true"`
	Name     string `json:"name"`
	Email    string `json:"email" required:"true" unique:"true"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role"`
}

// usecases
type UserUsecase interface {
	CreateUser(user User) error
	GetUserByEmail(email string) (User, error)
	Login(email, password string) (string, error)
	Promote(email string) error
}

type TaskUsecase interface {
	GetTasks() []Task
	GetTaskByID(id string) (Task, error)
	CreateTask(Task Task) error
	UpdateTask(id string, task Task) error
	DeleteTask(id string) error
}

// Repository
type UserRepository interface {
	CreateUser(user User) error
	GetUserByEmail(email string) (User, error)
	UpdateUser(email string, update User) error
}

type TaskRepository interface {
	GetTasks() []Task
	GetTaskByID(email string) (Task, error)
	CreateTask(Task Task) error
	DeleteTask(id string) error
	UpdateTask(id string, task Task) error
}

// services
type PasswordService interface {
	HashPassword(password string) (string, error)
	ComparePassword(pass1, pass2 string) error
}

type JWTService interface {
	GenerateToken(user User) (string, error)
	ValidateToken(role string,headers string) error
}


