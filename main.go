package main

import (
	"fmt"
	"log"
	"net/http"
	"we-web-api/book"
	"we-web-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=postgres password=yudistirar626 dbname=web-api port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection Failed")
	}

	fmt.Println("DB Connection Success")

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	// SERVICE

	// ==========
	// CREATE DATA REPOSITORY
	// ==========

	bookRequest := book.BookRequest{
		Title: "Test Service Terbaru",
		Price : "20000",
	}
	bookService.Create(bookRequest)


	
	// REPOSITORY

	// ==========
	// FIND ALL DATA REPOSITORY
	// ==========
	books, err := bookRepository.FindAll()

	for _, book := range books {
		fmt.Println("Title : ", book.Title)
	}

	// ==========
	// FIND BY ID REPOSITORY
	// ==========

	// book, err := bookRepository.FindById(3)

	// fmt.Println("Title : ", book.Title)

	// ==========
	// CREATE DATA REPOSITORY
	// ==========
	// book := book.Book {
	// 	Title : "C#",
	// 	Description : "C# is a programming language",
	// 	Price : 250000,
	// 	Rating : 5,
	// }

	// bookRepository.Create(book)


	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"name"	:	"Yudistira Rivaldi",
				"role"	:   "Software Engineer",
			})
	})

	// ROUTE BELAJAR
	// v1.GET("/hello", bookHandler.HelloHandler )
	// v1.GET("/books/:id", bookHandler.BooksHandler)
	// v1.GET("/book/:id/:title", bookHandler.BookHandler)
	// v1.GET("/query", bookHandler.QueryHandler)
		
	v1.POST("/bookspost", bookHandler.CreateBooksHandler)
	v1.GET("/books", bookHandler.GetBooksHandler)
	v1.GET("/book/:id", bookHandler.GetBookById)
	v1.PUT("/book/:id", bookHandler.UpdateBooksHandler)
	v1.DELETE("/book/:id", bookHandler.DeleteBooksHandler)
	

	router.Run(":8080")
}

// ===========
	// CREATE DATA
	// ===========

	// book := book.Book{}
	// book.Title = "Javascript"
	// book.Description = "Javascript is a programming language"
	// book.Price = 20
	// book.Rating = 8

	// err = db.Create(&book).Error

	// if err != nil {
	// 	fmt.Println("Error detected when create data")
	// }

	// ===============================================
	// READ GET DATA
	// ===============================================

	// book := book.Book{}

	// err = db.Debug().Where("id = ?", 2).First(&book).Error
	// if err != nil {
	// 	fmt.Println("Error detected when get data")
	// }

	// ===============================================
	// DELETE DATA
	// ===============================================

	// err = db.Delete(&book).Error
	// if err != nil {
	// 	fmt.Println("Error detected when delete data")
	// }


	// ===============================================
	// UPDATE DATA
	// ===============================================
	
	// book.Title = "Golang Book"

	// err = db.Save(&book).Error
	// if err != nil {
	// 	fmt.Println("Error detected when update data")
	// }




