package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/piipiets/go-library/handler"
	"github.com/piipiets/go-library/middleware"
)

func SetupRoutes(
	router *gin.Engine,
	authHandler *handler.AuthHandler,
	categoryHandler *handler.CategoryHandler,
) {
	router.POST("/login", authHandler.Login)

	// Protected routes
	protected := router.Group("/api")
	protected.Use(middleware.JWTAuth())

	protected.POST("/categories", categoryHandler.CreateCategory)

}
