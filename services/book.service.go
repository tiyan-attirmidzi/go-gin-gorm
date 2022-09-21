package services

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/tiyan-attirmidzi/go-gin-gorm/dto"
	"github.com/tiyan-attirmidzi/go-gin-gorm/entities"
	"github.com/tiyan-attirmidzi/go-gin-gorm/repositories"
)

type BookService interface {
	Index() []entities.Book
	Show(bookID uint64) entities.Book
	Store(book dto.BookCreate) entities.Book
	Update(book dto.BookUpdate) entities.Book
	Delete(book entities.Book) entities.Book
	IsAllowedToEdit(userID string, bookID uint64) bool
}

type bookService struct {
	bookRepository repositories.BookRepository
}

func NewBookService(bookRepository repositories.BookRepository) BookService {
	return &bookService{
		bookRepository: bookRepository,
	}
}

func (s *bookService) Index() []entities.Book {
	return s.bookRepository.Index()
}

func (s *bookService) Show(bookID uint64) entities.Book {
	return s.bookRepository.Show(bookID)
}

func (s *bookService) Store(book dto.BookCreate) entities.Book {
	data := entities.Book{}
	err := smapping.FillStruct(&data, smapping.MapFields(&book))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	return s.bookRepository.Store(data)
}

func (s *bookService) Update(book dto.BookUpdate) entities.Book {
	data := entities.Book{}
	err := smapping.FillStruct(&data, smapping.MapFields(&book))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	return s.bookRepository.Update(data)
}

func (s *bookService) Delete(book entities.Book) entities.Book {
	return s.bookRepository.Delete(book)
}

func (s *bookService) IsAllowedToEdit(userID string, bookID uint64) bool {
	book := s.bookRepository.Show(bookID)
	id := fmt.Sprintf("%v", book.UserID)
	return userID == id
}
