package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)

	router.GET("/hello", helloHandler)

	router.Run()

}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"name" : "Yudistira Rivaldi",
		"bio" : "I'm a web developer",
	})
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"message" : "Hello World",
		"quotes" : "Keep learning and you will be a backend developer",
	})
}