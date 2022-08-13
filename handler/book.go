package handler

import (
	"fmt"
	"net/http"
	"we-web-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
		"from":    "Yudistira Rivaldi",
	})
}

func BooksHandler(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"message": "Hello World",
	})
}

func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

func BookHandler(c *gin.Context) {

	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})

}

func PostBooksHandler(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

	if err != nil {

		// make multiple error
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})

	}

	c.JSON(http.StatusOK, gin.H{
		"title": bookRequest.Title,
		"price": bookRequest.Price,
	})

}