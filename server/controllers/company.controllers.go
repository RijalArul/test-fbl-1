package controllers

import (
	"net/http"
	"test-fbl-1/server/helpers"
	"test-fbl-1/server/services"
	"test-fbl-1/server/webs"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type CompanyController interface {
	Create(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

type CompanyControllerImpl struct {
	companyService services.CompanyService
}

func NewCompanyController(CompanyService services.CompanyService) CompanyController {
	return &CompanyControllerImpl{companyService: CompanyService}
}

func (c *CompanyControllerImpl) Create(ctx *gin.Context) {
	var inputCompany webs.CompanyDTO
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	contentType := helpers.GetContentType(ctx)

	if contentType == appJSON {
		ctx.ShouldBindJSON(&inputCompany)
	} else {
		ctx.ShouldBind(&inputCompany)
	}

	createCompany, err := c.companyService.Create(inputCompany, userID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ResponseSuccess(ctx, http.StatusCreated, createCompany)
}

func (c *CompanyControllerImpl) FindAll(ctx *gin.Context) {
	companies, err := c.companyService.FindAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ResponseSuccess(ctx, http.StatusOK, companies)
}
