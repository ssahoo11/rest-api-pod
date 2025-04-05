package middleware

import (
	"net/http"
	"server/utils"

	"github.com/gin-gonic/gin"
)

func Authenticate(cxt *gin.Context) {
	token := cxt.Request.Header.Get("Authorization")

	if token == "" {
		cxt.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		cxt.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	cxt.Set("userId", userId)
	cxt.Next()
}
