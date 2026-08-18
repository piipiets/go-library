package service

import (
	"errors"
	"time"

	"github.com/piipiets/go-library/model"
	"github.com/piipiets/go-library/model/request"
	"github.com/piipiets/go-library/repository"
)

var ErrBookNotFound = errors.New("book not found")

type BookService struct {
	repository *repository.BookRepository
}

func NewBookService(repository *repository.BookRepository) *BookService {
	return &BookService{
		repository: repository,
	}
}

func (s *BookService) CreateBook(bookRequest request.BookRequest, username string) error {
	now := time.Now()

	var book model.Book

	if bookRequest.ReleaseYear < 1980 && bookRequest.ReleaseYear > 2024 {
		return errors.New("Release year is not valid")
	}

	book.Title = bookRequest.Title
	book.CategoryId = bookRequest.CategoryId
	book.Description = bookRequest.Description
	book.ImageUrl = bookRequest.ImageUrl
	book.ReleaseYear = bookRequest.ReleaseYear
	book.Price = bookRequest.Price
	book.TotalPage = bookRequest.TotalPage
	book.CreatedAt = now
	book.CreatedBy = username
	book.ModifiedAt = now
	book.ModifiedBy = username

	if bookRequest.TotalPage > 100 {
		book.Thickness = "Tebal"
	} else {
		book.Thickness = "Tipis"
	}

	err := s.repository.CreateBook(book, username)
	if err != nil {
		return err
	}

	return nil
}
