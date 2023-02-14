package controllers

import (
	"net/http"
	"test-fbl-1/server/helpers"
	"test-fbl-1/server/services"
	"test-fbl-1/server/webs"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type TransactionController interface {
	Create(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

type TransactionControllerImpl struct {
	transactionService services.TransactionService
}

func NewTransactionController(TransactionService services.TransactionService) TransactionController {
	return &TransactionControllerImpl{transactionService: TransactionService}
}

func (c *TransactionControllerImpl) Create(ctx *gin.Context) {
	var inputTransaction webs.TransactionDTO
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	contentType := helpers.GetContentType(ctx)

	if contentType == appJSON {
		ctx.ShouldBindJSON(&inputTransaction)
	} else {
		ctx.ShouldBind(&inputTransaction)
	}

	createTransaction, err := c.transactionService.Create(inputTransaction, userID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ResponseSuccess(ctx, http.StatusCreated, createTransaction)
}

func (c *TransactionControllerImpl) FindAll(ctx *gin.Context) {
	transactions, err := c.transactionService.FindAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	ResponseSuccess(ctx, http.StatusOK, transactions)
}
