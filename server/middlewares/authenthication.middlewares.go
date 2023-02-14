package middlewares

import (
	"net/http"
	"test-fbl-1/server/helpers"

	"github.com/gin-gonic/gin"
)

func Authenthication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := helpers.VerifyToken(ctx)
		_ = verifyToken

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthenticated",
			})
			return

		}
		ctx.Set("userData", verifyToken)
		ctx.Next()
	}
}
