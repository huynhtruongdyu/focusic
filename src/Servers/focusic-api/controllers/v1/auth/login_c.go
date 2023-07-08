package v1

import "github.com/gin-gonic/gin"

func Login(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Login work!",
	})
}
