package v1

import "github.com/gin-gonic/gin"

func Logout(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Logout work!",
	})
}
