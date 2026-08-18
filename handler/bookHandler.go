package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/piipiets/go-library/model/request"
	"github.com/piipiets/go-library/model/response"
	"github.com/piipiets/go-library/service"
)

type BookHandler struct {
	service *service.BookService
}

func NewBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{
		service: service,
	}
}

func (h *BookHandler) CreateBook(c *gin.Context) {

	var bookRequest request.BookRequest
	// Parse JSON
	if err := c.ShouldBindJSON(&bookRequest); err != nil {
		c.JSON(http.StatusBadRequest, response.Error("Invalid request body", err.Error()))
		return
	}

	usernameValue, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.Error("Unauthorized", "username not found in token"))
		return
	}

	username, ok := usernameValue.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, response.Error("Unauthorized", "invalid username in token"))
		return
	}

	err := h.service.CreateBook(bookRequest, username)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error("Failed to create book", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, response.Success("Book created successfully", nil))
}
