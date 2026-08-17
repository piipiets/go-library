package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/piipiets/go-library/model/request"
	"github.com/piipiets/go-library/model/response"
	"github.com/piipiets/go-library/service"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var request request.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.Error("Invalid request body", err.Error()))
		return
	}

	token, err := h.authService.Login(
		request.Username,
		request.Password,
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, response.Error("Login failed", err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success("Login success", response.LoginResponse{
		Token: token,
	}))
}
