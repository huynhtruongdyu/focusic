package v1

import (
	authController "focusic-api/controllers/v1/auth"
	"github.com/gin-gonic/gin"
)

const AuthPrefix = "/auth"

func SetAuthRoutes(g *gin.RouterGroup) {
	g.GET(buildRoute(AuthPrefix, "login"), authController.Login)
	g.GET(buildRoute(AuthPrefix, "logout"), authController.Logout)
	g.POST(buildRoute(AuthPrefix, "signup"), authController.SignUp)
}
