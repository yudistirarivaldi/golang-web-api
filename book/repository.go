package book

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	Create(book Book) (Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repository {
	return repository{db}
}

func (r *repository) FindAll() ([]Book, error) {

	var books []Book
	err := r.db.Find(&books).Error
	
	return books, err
}

func (r *repository) FindById(ID int) (Book, error) {
	var books []Book

	err := r.db.whe
}