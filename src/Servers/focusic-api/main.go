package main

import (
	"fmt"
	"focusic-api/config"
	"focusic-api/database"
	v1 "focusic-api/routers/v1"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	router *gin.Engine
	port   string
)

func init() {
	appConfig := config.AppConfig
	if appConfig == nil {
		log.Fatalln("cannot.load.env")
	}
	database.ConnectSqlite()

	router = gin.Default()
	router.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, appConfig.AppName)
	})
	router.Static("/public", "/public")
	router.StaticFile("/favicon.ico", "./public/images/nuxt-logo.png")

	apiRoutes := router.Group("/api")
	version1 := apiRoutes.Group("/v1")
	v1.InitRoutes(version1)

	port = ":" + appConfig.Port
}

func main() {
	fmt.Println("Server Running on Port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
