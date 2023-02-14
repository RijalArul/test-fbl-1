package middlewares

import (
	"net/http"
	databases "test-fbl-1/server/db"
	"test-fbl-1/server/entities"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AdminAuthorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := databases.GetDB()
		user := entities.User{}
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		err := db.Model(user).Where("id = ?", userID).First(&user).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "You can't access"})
			return
		}

		if user.Role != "Admin" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "You can't access"})
			return
		}
		ctx.Next()
	}
}
