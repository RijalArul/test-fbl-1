package routers

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"test-fbl-1/server/controllers"
	databases "test-fbl-1/server/db"
	"test-fbl-1/server/entities"
	"test-fbl-1/server/middlewares"
	"test-fbl-1/server/repositories"
	"test-fbl-1/server/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type M map[string]interface{}

func handleDownload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	path := r.FormValue("path")
	fmt.Println(path)
	f, err := os.Open(path)
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	contentDisposition := fmt.Sprintf("attachment; filename=%s", f.Name())
	w.Header().Set("Content-Disposition", contentDisposition)

	if _, err := io.Copy(w, f); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Routes() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	DB := databases.GetDB()
	var transactions []entities.Transaction
	err := DB.Preload(clause.Associations).Model(transactions).Find(&transactions).Error
	file, err := os.Create("records.csv")
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(file)
	defer w.Flush()
	var data [][]string
	for _, record := range transactions {
		row := []string{strconv.Itoa(int(record.ID)), record.CompanyName, record.ProductName, strconv.Itoa(record.TotalQuantity), strconv.Itoa(record.TotalPrice), strconv.Itoa(record.Product.Stock)}
		data = append(data, row)
	}
	w.WriteAll(data)

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

	transactionRepository := repositories.NewTransactionRepository(DB)
	transactionService := services.NewTransactionService(transactionRepository, companyRepository, productRepository)
	transactionController := controllers.NewTransactionController(transactionService)
	transactionRouter := r.Group("/transactions")
	{
		transactionRouter.Use(middlewares.Authenthication())
		transactionRouter.POST("/", transactionController.Create)
		transactionRouter.GET("/", transactionController.FindAll)
	}
	http.HandleFunc("/download", handleDownload)
	r.Run()
	return r
}
