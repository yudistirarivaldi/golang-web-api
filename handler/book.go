package handler

import (
	"fmt"
	"net/http"
	"we-web-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
		"from":    "Yudistira Rivaldi",
	})
}

func (h *bookHandler) BooksHandler(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"message": "Hello World",
	})
}

func (h *bookHandler) QueryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

func (h *bookHandler) BookHandler(c *gin.Context) {

	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})

}

func (h *bookHandler) PostBooksHandler(c *gin.Context) {
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

	// call service
	book, err := h.bookService.Create(bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"errors" : err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data" : book,
	})

}