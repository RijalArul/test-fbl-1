package helpers

import (
	"errors"
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = "Kkbsbxhbdhbhbjb29u39e9823u9bub934y84f9g39fb9ub.kaskaksnasuhasuh"

func GenerateToken(id uint, username string, role string) string {
	claims := jwt.MapClaims{
		"id":       id,
		"username": username,
		"role":     role,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := parseToken.SignedString([]byte(secretKey))
	return signedToken
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	errReponse := errors.New("sign in to proceed")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")
	if !bearer {
		return nil, errReponse
	}
	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errReponse
		}
		return []byte(secretKey), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errReponse
	}
	log.Println(token.Claims.(jwt.MapClaims))
	return token.Claims.(jwt.MapClaims), nil
}
