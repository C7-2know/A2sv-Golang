package router

import (
	"task_manager/Delivery/controllers"
	domain "task_manager/Domain"
	service "task_manager/Infrastructures"
	repository "task_manager/Repositories"
	usecases "task_manager/Usecases"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetUp(db mongo.Database, gin *gin.Engine) {
	pass_service := service.NewPasswordService()
	jwt := service.NewJwtService()
	publicRoute := gin.Group("")
	NewSignUpRoute(pass_service,jwt,db, publicRoute)

	userRoute := gin.Group("")
	userRoute.Use(service.AuthMiddleware(""))
	
	NewUserRoute(pass_service, jwt, db, userRoute)

	adminRoute := gin.Group("")
	adminRoute.Use(service.AuthMiddleware("admin"))
	NewAdminRoute(pass_service,jwt,db, adminRoute)

}

func NewSignUpRoute(pass domain.PasswordService, jwt domain.JWTService, db mongo.Database, route *gin.RouterGroup) {
	repository := repository.NewUserRepository(db)

	uu := usecases.NewUserUsecase(repository, pass, jwt)
	controller := &controllers.UserController{User_usecase: uu}
	// controller :=&controllers.UserController{user_usecase:userUsecase,}
	route.POST("/register", controller.SignUp)
}

func NewUserRoute(pass domain.PasswordService, jwt domain.JWTService, db mongo.Database, route *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	tr := repository.NewTaskRepository(db)
	uu := usecases.NewUserUsecase(ur, pass, jwt)
	tu := usecases.NewTaskUsecase(tr)
	uc := &controllers.UserController{User_usecase: uu}
	tc := &controllers.TaskController{Task_usecase: tu}
	route.GET("/login", uc.LogIn)
	route.GET("/tasks", tc.GetTasks)
	route.GET("/tasks:id", tc.GetTask)
}
func NewAdminRoute(pass domain.PasswordService, jwt domain.JWTService,db mongo.Database, route *gin.RouterGroup) {
	repo := repository.NewUserRepository(db)
	tr := repository.NewTaskRepository(db)
	uu := usecases.NewUserUsecase(repo,pass,jwt)
	tu := usecases.NewTaskUsecase(tr)

	uc := &controllers.UserController{User_usecase: uu}
	tc := &controllers.TaskController{Task_usecase: tu}

	route.PUT("/promote/:email", uc.PromoteUser)
	route.POST("/tasks", tc.CreateTask)
	route.PUT("/tasks/:id", tc.UpdateTask)
	route.DELETE("/tasks", tc.DeleteTask)

}
