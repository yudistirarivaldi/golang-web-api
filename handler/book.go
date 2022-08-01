package handler

import (
	"fmt"
	"net/http"

	"golang-web-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Yudistira Rivaldi",
		"bio":  "I'm a web developer",
	})
}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
		"quotes":  "Keep learning and you will be a backend developer",
	})
}

func BooksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})

	// output nya http://localhost:8080/books/1

}

func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})

	// output nya http://localhost:8080/query?title=Bumi Manusia&price=15000
}

func PostBooksHandler(c *gin.Context) {
	var bookInput book.BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		
		}

		c.JSON(http.StatusBadRequest, gin.H {
			"errors" : errorMessages,
		
		})
		return
	}

	c.JSON(http.StatusOK, gin.H {
		"title" : bookInput.Title,
		"price" : bookInput.Price,
		//"subtitle" : bookInput.Subtitle,
	})


}