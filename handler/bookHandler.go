package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/piipiets/go-library/model/request"
	"github.com/piipiets/go-library/model/response"
	"github.com/piipiets/go-library/service"
)

type BookHandler struct {
	service service.BookService
}

func NewBookHandler(service service.BookService) *BookHandler {
	return &BookHandler{
		service: service,
	}
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var bookRequest request.BookRequest

	if err := c.ShouldBindJSON(&bookRequest); err != nil {
		c.JSON(
			http.StatusBadRequest,
			response.Error("Invalid request body", err.Error()),
		)
		return
	}

	usernameValue, exists := c.Get("username")
	if !exists {
		c.JSON(
			http.StatusUnauthorized,
			response.Error("Unauthorized", "username not found in token"),
		)
		return
	}

	username, ok := usernameValue.(string)
	if !ok {
		c.JSON(
			http.StatusUnauthorized,
			response.Error("Unauthorized", "invalid username in token"),
		)
		return
	}

	err := h.service.CreateBook(bookRequest, username)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			response.Error("Failed to create book", err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		response.Success("Book created successfully", nil),
	)
}

func (h *BookHandler) GetAllBooks(c *gin.Context) {
	books, err := h.service.GetAllBooks()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			response.Error("Failed to get books data", err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		response.Success("Books fetched successfully", books),
	)
}

func (h *BookHandler) GetBookById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			response.Error("Invalid book id", "book id must be a number"),
		)
		return
	}

	book, err := h.service.GetBookById(id)

	if err != nil {
		if err == service.ErrBookNotFound {
			c.JSON(
				http.StatusNotFound,
				response.Error("Book not found", err.Error()),
			)
			return
		}

		c.JSON(
			http.StatusInternalServerError,
			response.Error("Failed to get book data", err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		response.Success("Book fetched successfully", book),
	)
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			response.Error("Invalid book id", "book id must be a number"),
		)
		return
	}

	var bookRequest request.BookRequest

	if err := c.ShouldBindJSON(&bookRequest); err != nil {
		c.JSON(
			http.StatusBadRequest,
			response.Error("Invalid request body", err.Error()),
		)
		return
	}

	usernameValue, exists := c.Get("username")
	if !exists {
		c.JSON(
			http.StatusUnauthorized,
			response.Error("Unauthorized", "username not found in token"),
		)
		return
	}

	username, ok := usernameValue.(string)
	if !ok {
		c.JSON(
			http.StatusUnauthorized,
			response.Error("Unauthorized", "invalid username in token"),
		)
		return
	}

	err = h.service.UpdateBook(id, bookRequest, username)

	if err != nil {
		if err == service.ErrBookNotFound {
			c.JSON(
				http.StatusNotFound,
				response.Error("Book not found", err.Error()),
			)
			return
		}

		c.JSON(
			http.StatusBadRequest,
			response.Error("Failed to update book", err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		response.Success("Book updated successfully", nil),
	)
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			response.Error("Invalid book id", "book id must be a number"),
		)
		return
	}

	err = h.service.DeleteBook(id)

	if err != nil {
		if err == service.ErrBookNotFound {
			c.JSON(
				http.StatusNotFound,
				response.Error("Book not found", err.Error()),
			)
			return
		}

		c.JSON(
			http.StatusInternalServerError,
			response.Error("Failed to delete book", err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		response.Success("Book deleted successfully", nil),
	)
}
