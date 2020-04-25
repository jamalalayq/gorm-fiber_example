package book

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// Book refer to book
type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

// Repository generic repo.
type Repository interface {
	GetAll() []interface{}
	GetOne(id string) interface{}
	Save(item interface{}) interface{}
	Delete(id string) interface{}
}

// BooksRepository repo for books
type BooksRepository struct {
	DB *gorm.DB
}

// GetAll fetch all saved books.
func (br BooksRepository) GetAll() []Book {
	var books []Book
	br.DB.Find(&books)
	return books
}

// GetOne fetch book by id from database
func (br BooksRepository) GetOne(id string) (Book, error) {
	var book Book
	br.DB.Find(&book, id)
	if book.Title == "" {
		return Book{}, errors.New("can't find this book")
	}
	return book, nil
}

// Save add new book to database
func (br BooksRepository) Save(item Book) Book {
	var book Book = item
	br.DB.Create(&book)
	return book
}

// Delete a book by id from database
func (br BooksRepository) Delete(id string) error {
	var book Book
	br.DB.First(&book, id)
	if book.Title == "" {
		return errors.New("This book not found")
	}
	br.DB.Delete(&book)
	return nil
}
