package routes

import (
	"fmt"
	"net/http"
	"strings"

	"db-go.com/api/tokens"
	"github.com/gin-gonic/gin"
)

func validateAuth(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	if authHeader == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	headerToken := strings.SplitN(authHeader, " ", 2)
	if len(headerToken) != 2 || headerToken[0] != "Bearer" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	claims, err := tokens.ValidateToken(headerToken[1])
	fmt.Println(claims)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}
	context.Set("claims", claims)
	context.Next()
}
