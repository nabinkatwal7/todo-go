package middleware

import (
	"net/http"
	"todo-api/helper"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc{
	return func(context *gin.Context){
		err := helper.ValidateJWT(context)
		if err != nil{
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			context.Abort()
			return
		}
		context.Next()
	}
}