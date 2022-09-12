package repositories

import (
	"github.com/tiyan-attirmidzi/go-rest-api/entities"
	"gorm.io/gorm"
)

type BookRepository interface {
	Index() []entities.Book
	Show(bookID uint64) entities.Book
	Store(book entities.Book) entities.Book
	Update(book entities.Book) entities.Book
	Delete(book entities.Book) entities.Book
}

type bookConnection struct {
	connection *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookConnection{
		connection: db,
	}
}

func (db *bookConnection) Index() []entities.Book {
	var books []entities.Book
	db.connection.Preload("User").Find(&books)
	return books
}

func (db *bookConnection) Show(bookID uint64) entities.Book {
	var book entities.Book
	db.connection.Preload("User").Find(&book, bookID)
	return book
}

func (db *bookConnection) Store(book entities.Book) entities.Book {
	db.connection.Save(&book)
	db.connection.Preload("User").Find(&book)
	return book
}

func (db *bookConnection) Update(book entities.Book) entities.Book {
	db.connection.Save(&book)
	db.connection.Preload("User").Find(&book)
	return book

}

func (db *bookConnection) Delete(book entities.Book) entities.Book {
	db.connection.Delete(&book)
	return book
}
