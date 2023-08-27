package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/aalsa16/golang-microservices/utils"
	"github.com/gin-gonic/gin"
)

type AuthBody struct {
	Uuid string `json:"uuid"`
}

func getTokenFromHeader(reqToken string) (string, error) {
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) != 2 {
		return "", errors.New("Invalid Authorization header format")
	}
	idToken := splitToken[1]
	return idToken, nil
}

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		val, provided := context.GetQuery("uuid")
		if !provided || len(val) == 0 {
			context.JSON(http.StatusBadRequest, gin.H{"error": "uuid value must be provided"})
			context.Abort()
			return
		}

		reqToken := context.GetHeader("Authorization")
		idToken, err := getTokenFromHeader(reqToken)

		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}

		if idToken == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}

		err = utils.ValidateToken(idToken, val)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}

		context.Next()
	}
}
