package v1

import (
	"github.com/gin-gonic/gin"
	"path"
)

// @BasePath /api/v1
// PingExample godoc
func InitRoutes(g *gin.RouterGroup) {
	SetAuthRoutes(g)
	SetUsersRoutes(g)
}

func buildRoute(prefix, endpointPath string) string {
	return path.Join(prefix, endpointPath)
}
