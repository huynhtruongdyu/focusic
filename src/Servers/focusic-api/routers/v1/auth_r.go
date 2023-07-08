package v1

import (
	authController "focusic-api/controllers/v1/auth"
	"github.com/gin-gonic/gin"
)

const prefix = "/auth"

func SetAuthRoutes(g *gin.RouterGroup) {
	g.GET(prefix+"/login", authController.Login)
	g.GET(prefix+"/logout", authController.Logout)
	g.POST(prefix+"/signup", authController.SignUp)
}
