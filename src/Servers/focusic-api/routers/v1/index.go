package v1

import (
	"github.com/gin-gonic/gin"
	"path"
)

func InitRoutes(g *gin.RouterGroup) {
	SetAuthRoutes(g)
}

func buildRoute(prefix, endpointPath string) string {
	return path.Join(prefix, endpointPath)
}
