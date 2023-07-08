package v1

import (
	"focusic-api/config"
	model "focusic-api/models"
	"focusic-api/models/domains"
	"focusic-api/models/dto"
	"focusic-api/repositories"
	"github.com/gin-gonic/gin"
)

var validator = config.AppValidator

func SignUp(ctx *gin.Context) {
	var userCreateModel dto.UserCreateRequestModel

	if err := ctx.Bind(&userCreateModel); err != nil {
		ctx.JSON(200, model.ApiFailedResult("invalid.data"))
		return
	}

	if validationErr := validator.Struct(&userCreateModel); validationErr != nil {
		ctx.JSON(200, model.ApiFailedResult("invalid.data"))
		return
	}

	userRepo := repositories.GetUserRepository()

	user, err := userRepo.Create(domains.User{
		Name: userCreateModel.Name,
	})
	if err != nil {
		ctx.JSON(200, model.ApiFailedResult("create.user.failed"))
	}
	response := model.ApiSuccessResult(user)
	ctx.JSON(200, response)
}
