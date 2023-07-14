package main

import (
	"fmt"
	"focusic-api/config"
	"focusic-api/database"
	_ "focusic-api/docs"
	v1 "focusic-api/routers/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, appConfig.AppName)
	})
	router.Static("/public", "/public")
	router.StaticFile("/favicon.ico", "./public/images/logo-icon.png")

	apiRoutes := router.Group("/api")
	version1 := apiRoutes.Group("/v1")

	v1.InitRoutes(version1)

	port = ":" + appConfig.Port
}

// @title           Focusic API
// @version         1.0
// @description     focusic api server
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @securityDefinitions.apikey  BearerAuth
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	fmt.Println("Server Running on Port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
