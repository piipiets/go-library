package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/piipiets/go-library/model/request"
	"github.com/piipiets/go-library/model/response"
	"github.com/piipiets/go-library/service"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		service: service,
	}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {

	var categoryRequest request.CategoryRequest
	// Parse JSON
	if err := c.ShouldBindJSON(&categoryRequest); err != nil {
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

	category, err := h.service.CreateCategory(categoryRequest, username)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error("Failed to create category", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, response.Success("Category created successfully", category))
}

func (h *CategoryHandler) GetAllCategory(c *gin.Context) {
	categories, err := h.service.GetAllCategories()

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error("Failed to get categories data", err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success("Categories fetched successfully", categories))
}

func (h *CategoryHandler) GetCategoryById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error("Invalid category ID", err.Error()))
		return
	}

	category, err := h.service.GetCategoryById(id)

	if err != nil {
		if errors.Is(err, service.ErrCategoryNotFound) {
			c.JSON(http.StatusNotFound, response.Error("Category Not Found", err.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, response.Error("Failed to get category data", err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success("Category fetched successfully", category))
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error("Invalid category ID", err.Error()))
		return
	}

	err2 := h.service.DeleteCategory(id)

	if err2 != nil {
		if errors.Is(err2, service.ErrCategoryNotFound) {
			c.JSON(http.StatusNotFound, response.Error("Category Not Found", err2.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, response.Error("Failed to delete category data", err2.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success("Category deleted successfully", nil))
}

func (h *CategoryHandler) GetBooksByCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error("Invalid category ID", err.Error()))
		return
	}

	books, err := h.service.GetBooksByCategory(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error("Failed to get books data", err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success("Books fetched successfully", books))
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	id, errParam := strconv.Atoi(c.Param("id"))
	if errParam != nil {
		c.JSON(http.StatusBadRequest, response.Error("Invalid category ID", errParam.Error()))
		return
	}

	var categoryRequest request.CategoryRequest
	// Parse JSON
	if errBody := c.ShouldBindJSON(&categoryRequest); errBody != nil {
		c.JSON(http.StatusBadRequest, response.Error("Invalid request body", errBody.Error()))
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

	err := h.service.UpdateCategory(id, categoryRequest, username)

	if err != nil {
		if errors.Is(err, service.ErrCategoryNotFound) {
			c.JSON(http.StatusNotFound, response.Error("Category Not Found", err.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, response.Error("Failed to update category data", err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success("Category updated successfully", nil))
}
