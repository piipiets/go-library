package service

import (
	"database/sql"
	"errors"
	"time"

	"github.com/piipiets/go-library/model"
	"github.com/piipiets/go-library/model/request"
	"github.com/piipiets/go-library/model/response"
	"github.com/piipiets/go-library/repository"
)

var ErrBookNotFound = errors.New("book not found")

type BookService interface {
	CreateBook(bookRequest request.BookRequest, username string) error
	GetAllBooks() ([]response.BookResponse, error)
	GetBookById(id int) (response.BookResponse, error)
	UpdateBook(id int, req request.BookRequest, username string) error
	DeleteBook(id int) error
}
type bookService struct {
	repository repository.Repository
}

func NewBookService(repository repository.Repository) BookService {
	return &bookService{
		repository: repository,
	}
}

func (s *bookService) CreateBook(bookRequest request.BookRequest, username string) error {
	var book model.Book

	if bookRequest.ReleaseYear < 1980 || bookRequest.ReleaseYear > 2024 {
		return errors.New("Release year is not valid")
	}

	book = ToBookModel(bookRequest, username)

	err := s.repository.CreateBook(book)
	if err != nil {
		return err
	}

	return nil
}

func (s *bookService) GetAllBooks() ([]response.BookResponse, error) {
	books, err := s.repository.GetAllBooks()
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (s *bookService) GetBookById(id int) (response.BookResponse, error) {
	book, err := s.repository.GetBookById(id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.BookResponse{}, ErrBookNotFound
		}

		return response.BookResponse{}, err
	}

	return book, nil
}

func (s *bookService) UpdateBook(id int, req request.BookRequest, username string) error {
	var book model.Book

	if req.ReleaseYear < 1980 || req.ReleaseYear > 2024 {
		return errors.New("Release year is not valid")
	}

	book = ToBookModel(req, username)

	err := s.repository.UpdateBook(id, book)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrCategoryNotFound
		}

		return err
	}

	return nil
}

func ToBookModel(bookRequest request.BookRequest, username string) (book model.Book) {
	now := time.Now()
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

	return book
}

func (s *bookService) DeleteBook(id int) error {
	err := s.repository.DeleteBook(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrBookNotFound
		}

		return err
	}

	return nil
}
