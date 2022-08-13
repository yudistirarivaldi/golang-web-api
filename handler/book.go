package handler

import (
	"fmt"
	"net/http"
	"strconv"
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

func (h *bookHandler) GetBooksHandler(c *gin.Context) {
	books, err := h.bookService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"errors" : err,
		})
		return
	}

	var booksResponse []book.BookResponse

	for _, b := range books {
		bookResponse := book.BookResponse {
			ID: 		 b.ID,
			Title: 	     b.Title,
			Description: b.Description,
			Price: 		 b.Price,
			Rating: 	 b.Rating,
		}

		booksResponse = append(booksResponse, bookResponse)

	}

	c.JSON(http.StatusOK, gin.H {
		"data" : booksResponse,
	})

}

func (h *bookHandler) GetBookById(c *gin.Context) {
	idString := c.Param("id") //karena id ini adalah string jadi harus di pindah ke int
	id, _ := strconv.Atoi(idString)
	
	b, err := h.bookService.FindById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"errors" : err,
		})
		return
	}

	// agar bawaan struct book berubah menjadi book response
	// cara 1
	// bookResponse := book.BookResponse {
	// 	ID: 		 b.ID,
	// 	Title: 	     b.Title,
	// 	Description: b.Description,
	// 	Price: 		 b.Price,
	// 	Rating: 	 b.Rating,
	// }

	// cara 2
	bookResponse := convertToBookResponse(b)


	c.JSON(http.StatusOK, gin.H {
		"data" : bookResponse,
	})

}

func (h *bookHandler) CreateBooksHandler(c *gin.Context) {
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
		"data" : convertToBookResponse(book),
	})
}

func (h *bookHandler) UpdateBooksHandler(c *gin.Context) {
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

	idString := c.Param("id") //karena id ini adalah string jadi harus di pindah ke int
	id, _ := strconv.Atoi(idString)

	// call service
	book, err := h.bookService.Update(id, bookRequest)

	

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"errors" : err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data" : convertToBookResponse(book),
	})
}

func (h *bookHandler) DeleteBooksHandler(c *gin.Context) {
	idString := c.Param("id") //karena id ini adalah string jadi harus di pindah ke int
	id, _ := strconv.Atoi(idString)
	
	b, err := h.bookService.Delete(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"errors" : err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H {
		"data" : convertToBookResponse(b),
	})

}



// private function convert struct book to json format book
func convertToBookResponse(b book.Book) book.BookResponse {
	return book.BookResponse {
		ID: 		 b.ID,
		Title: 	     b.Title,
		Description: b.Description,
		Price: 		 b.Price,
		Rating: 	 b.Rating,
	}
}




















// HANDLER BUAT BELAJAR

// func (h *bookHandler) HelloHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Hello World",
// 		"from":    "Yudistira Rivaldi",
// 	})
// }

// func (h *bookHandler) BooksHandler(c *gin.Context) {
// 	id := c.Param("id")

// 	c.JSON(http.StatusOK, gin.H{
// 		"id":      id,
// 		"message": "Hello World",
// 	})
// }

// func (h *bookHandler) QueryHandler(c *gin.Context) {
// 	title := c.Query("title")
// 	price := c.Query("price")

// 	c.JSON(http.StatusOK, gin.H{
// 		"title": title,
// 		"price": price,
// 	})
// }

// func (h *bookHandler) BookHandler(c *gin.Context) {

// 	id := c.Param("id")
// 	title := c.Param("title")

// 	c.JSON(http.StatusOK, gin.H{
// 		"id":    id,
// 		"title": title,
// 	})

// }