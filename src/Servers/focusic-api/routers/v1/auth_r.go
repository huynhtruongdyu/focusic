package v1

import (
	"focusic-api/controllers/v1/auth"
	"github.com/gin-gonic/gin"
)

const prefix = "/auth"

func SetAuthRoutes(g *gin.RouterGroup) {
	g.GET(prefix+"/login", auth.Login)
	g.GET(prefix+"/logout", auth.Logout)
}
