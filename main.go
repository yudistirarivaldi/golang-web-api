package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"name"	:	"Yudistira Rivaldi",
				"role"	:   "Software Engineer",
			})
	})
	v1.GET("/hello", helloHandler )
	v1.GET("/books/:id", booksHandler)
	v1.GET("/book/:id/:title", bookHandler)
	v1.GET("/query", queryHandler)
		
	v1.POST("/bookspost", postBooksHandler)

	router.Run()

}
func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
		"from": "Yudistira Rivaldi",
	})
}

func booksHandler(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H {
		"id": id,
		"message": "Hello World",
	})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H {
		"title" : title,
		"price" : price,
	})
}

func bookHandler(c *gin.Context) {

	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H {
		"id"	: id,
		"title" : title,
	})

}

type BookInput struct {
	Title string `json:"title" binding:"required" `
	Price json.Number `json:"price" binding:"required,number"`
}

func postBooksHandler(c *gin.Context) {
	var bookInput BookInput

	err := c.ShouldBindJSON(&bookInput)

	if err != nil {

		// make multiple error
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		return
		} 

			c.JSON(http.StatusBadRequest, gin.H {
				"error" : errorMessages,
			})

	}

	c.JSON(http.StatusOK, gin.H {
		"title" : bookInput.Title,
		"price" : bookInput.Price,

	})

}
