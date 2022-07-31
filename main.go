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

	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)
	router.GET("/books/:id/:title", booksHandler)
	router.GET("/query", queryHandler)

	router.POST("/books", postBooksHandler)

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

func booksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H {
		"id" : id,
		"title" : title,
	})

	// output nya http://localhost:8080/books/1

}

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H {
		"title" : title,
		"price"	: price,
	})

	// output nya http://localhost:8080/query?title=Bumi Manusia&price=15000
}

type BookInput struct {
	Title string 		`json:"title" binding:"required"`
	Price json.Number 	`json:"price" binding:"required,number"`

	//Subtitle string // if you want call subtitle with _ example sub_title you can use `json:"sub_title"`
}

func postBooksHandler(c *gin.Context) {
	var bookInput BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errorMessage)
			return
		}
		
	}

	c.JSON(http.StatusOK, gin.H {
		"title" : bookInput.Title,
		"price" : bookInput.Price,
		//"subtitle" : bookInput.Subtitle,
	})


}