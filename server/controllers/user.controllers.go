package controllers

import (
	"net/http"
	"test-fbl-1/server/helpers"
	"test-fbl-1/server/services"
	"test-fbl-1/server/webs"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func ResponseSuccess(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.JSON(statusCode, gin.H{
		"data": data,
	})
}

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type UserControllerImpl struct {
	userService services.UserService
}

func NewUserController(UserService services.UserService) UserController {
	return &UserControllerImpl{userService: UserService}
}

func (c *UserControllerImpl) Register(ctx *gin.Context) {
	var registerDTO webs.RegisterDTO
	contentType := helpers.GetContentType(ctx)

	if contentType == appJSON {
		ctx.ShouldBindJSON(&registerDTO)
	} else {
		ctx.ShouldBind(&registerDTO)
	}

	createUser, err := c.userService.Register(registerDTO)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ResponseSuccess(ctx, http.StatusCreated, createUser)
}

func (c *UserControllerImpl) Login(ctx *gin.Context) {
	var inputLogin webs.LoginDTO
	contentType := helpers.GetContentType(ctx)

	if contentType == appJSON {
		ctx.ShouldBindJSON(&inputLogin)
	} else {
		ctx.ShouldBind(&inputLogin)
	}

	loginUser, err := c.userService.Login(inputLogin)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	checkPass := helpers.ComparePass([]byte(loginUser.Password), []byte(inputLogin.Password))

	if checkPass == false {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Password"})
		return
	}

	genToken := helpers.GenerateToken(loginUser.ID, loginUser.Username, loginUser.Role)
	ResponseSuccess(ctx, http.StatusOK, genToken)
}
