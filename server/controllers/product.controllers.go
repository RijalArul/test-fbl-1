package controllers

import (
	"net/http"
	"test-fbl-1/server/helpers"
	"test-fbl-1/server/services"
	"test-fbl-1/server/webs"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type ProductController interface {
	Create(ctx *gin.Context)
}

type ProductControllerImpl struct {
	productService services.ProductService
}

func NewProductController(ProductService services.ProductService) ProductController {
	return &ProductControllerImpl{productService: ProductService}
}

func (c *ProductControllerImpl) Create(ctx *gin.Context) {
	var inputProduct webs.ProductDTO
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	contentType := helpers.GetContentType(ctx)

	if contentType == appJSON {
		ctx.ShouldBindJSON(&inputProduct)
	} else {
		ctx.ShouldBind(&inputProduct)
	}

	createProduct, err := c.productService.Create(inputProduct, userID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ResponseSuccess(ctx, http.StatusCreated, createProduct)
}
