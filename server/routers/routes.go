package routers

import (
	"test-fbl-1/server/controllers"
	databases "test-fbl-1/server/db"
	"test-fbl-1/server/middlewares"
	"test-fbl-1/server/repositories"
	"test-fbl-1/server/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	DB := databases.GetDB()

	userRepository := repositories.NewUserRepository(DB)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", userController.Register)
		userRouter.POST("/login", userController.Login)
	}

	companyRepository := repositories.NewCompanyRepository(DB)
	companyService := services.NewCompanyService(companyRepository)
	companyController := controllers.NewCompanyController(companyService)
	companyRouter := r.Group("/companies")

	{

		companyRouter.GET("/", companyController.FindAll)
		companyRouter.Use(middlewares.Authenthication())
		companyRouter.Use(middlewares.AdminAuthorize())
		companyRouter.POST("/", companyController.Create)
	}

	productRepository := repositories.NewProductRepository(DB)
	productService := services.NewProductService(productRepository, companyRepository)
	productController := controllers.NewProductController(productService)
	productRouter := r.Group("/products")

	{
		productRouter.Use(middlewares.Authenthication())
		productRouter.Use(middlewares.AdminAuthorize())
		productRouter.POST("/", productController.Create)
	}
	r.Run()
	return r
}
