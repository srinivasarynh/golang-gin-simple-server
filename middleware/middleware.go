package middleware

import "github.com/gin-gonic/gin"

func Authentication(ctx *gin.Context) {
	if !(ctx.Request.Header.Get("Token") == "auth") {
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "Auth Token is not present",
		})

		return
	}

	ctx.Next()
}
