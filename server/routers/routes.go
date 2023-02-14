package routers

import (
	"test-fbl-1/server/controllers"
	databases "test-fbl-1/server/db"
	"test-fbl-1/server/repositories"
	"test-fbl-1/server/services"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()

	DB := databases.GetDB()

	userRepository := repositories.NewUserRepository(DB)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", userController.Register)
		userRouter.POST("/login", userController.Login)
	}
	r.Run()
	return r
}
