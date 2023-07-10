package v1

import (
	"focusic-api/repositories"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(ctx *gin.Context) {
	userRepo := repositories.NewUserRepository()
	var users = userRepo.FindAll()
	ctx.JSON(200, users)
}
