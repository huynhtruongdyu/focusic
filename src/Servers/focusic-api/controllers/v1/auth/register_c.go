package v1

import (
	"focusic-api/config"
	"focusic-api/models/domains"
	"focusic-api/models/dto"
	"focusic-api/repositories"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var userCreateModel dto.UserCreateRequestModel
	err := ctx.BindJSON(&userCreateModel)
	if err != nil {
		ctx.JSON(404, gin.H{
			"errors": err.Error(),
		})
	}
	appValidator := config.AppValidator
	if err = appValidator.Struct(userCreateModel); err != nil {
		ctx.JSON(400, gin.H{
			"errors": err.Error(),
		})
	}
	user := domains.User{Name: userCreateModel.Name}
	userRepo := repositories.NewUserRepository()
	newUser, err := userRepo.Create(user)
	if err != nil {
		ctx.JSON(404, gin.H{
			"errors": err.Error(),
		})
	}

	ctx.JSON(200, newUser)
}
