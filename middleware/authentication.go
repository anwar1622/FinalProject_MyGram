package middleware

import (
	"MyGram/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authentication() gin.HandlerFunc {
	return func(context *gin.Context) {
		verifyToken, err := helpers.VerifyToken(context)
		_ = verifyToken

		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": err.Error(),
			})
			return
		}
		context.Set("userData", verifyToken)
		context.Next()
	}
}
