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

	// bookRepository := book.Repository{db}

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"name"	:	"Yudistira Rivaldi",
				"role"	:   "Software Engineer",
			})
	})
	v1.GET("/hello", handler.HelloHandler )
	v1.GET("/books/:id", handler.BooksHandler)
	v1.GET("/book/:id/:title", handler.BookHandler)
	v1.GET("/query", handler.QueryHandler)
		
	v1.POST("/bookspost", handler.PostBooksHandler)

	router.Run()
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




