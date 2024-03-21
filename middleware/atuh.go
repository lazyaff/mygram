package middleware

import (
	"final-project/helpers"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := helpers.ValidateToken(ctx)
		if err != nil {
			ctx.JSON(401, gin.H{
				"status":  "error",
				"code": "401",
				"message": err.Error(),
			})
			return
		}
		ctx.Set("userData", verifyToken)
		ctx.Next()
	}
}