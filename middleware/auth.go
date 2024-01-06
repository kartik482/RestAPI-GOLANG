package middleware

import (
	"net/http"

	"example.com/api/utils"
	"github.com/gin-gonic/gin"
)

func Authentication(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not provided"})
		return
	}
	userid, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	context.Set("userid", userid)
	context.Next()
}
