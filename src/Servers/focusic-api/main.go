package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "hello world")
	})
	log.Fatal(r.Run())
}
