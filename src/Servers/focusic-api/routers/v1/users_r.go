package v1

import (
	usersController "focusic-api/controllers/v1/users"
	"github.com/gin-gonic/gin"
)

const UsersPrefix = "/users"

func SetUsersRoutes(g *gin.RouterGroup) {
	g.GET(buildRoute(UsersPrefix, ""), usersController.GetAllUsers)
}
