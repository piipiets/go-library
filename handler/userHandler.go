package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/piipiets/go-library/model/request"
	"github.com/piipiets/go-library/model/response"
	"github.com/piipiets/go-library/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {

	var userReq request.AddUserRequest

	// Parse JSON
	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, response.Error("Invalid request body", err.Error()))
		return
	}

	user, err := h.service.CreateUser(userReq)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error("Failed to create user", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, response.Success("User created successfully", user))
}
