package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/piipiets/go-library/model/request"
	"github.com/piipiets/go-library/service"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHancder(service *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		service: service,
	}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {

	var categoryRequest request.CategoryRequest
	// Parse JSON
	if err := c.ShouldBindJSON(&categoryRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	username, _ := c.Get("username")

	category, err := h.service.CreateCategory(categoryRequest, username.(string))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, category)
}

func (h *CategoryHandler) GetAllCategory(c *gin.Context) {
	categories, err := h.service.GetAllCategories()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get categories data",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, categories)
}
