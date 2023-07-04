package main

import (
	"fmt"
	"focusic-api/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	appConfig := config.GetConfig()
	port := fmt.Sprintf(":%s", appConfig.Port)
	r := gin.Default()

	r.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, appConfig.AppName)
	})
	log.Fatal(r.Run(port))
}
