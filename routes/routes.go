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
	userHandler *handler.UserHandler,
) {
	router.POST("/login", authHandler.Login)
	router.POST("/user/add", userHandler.CreateUser)

	// Protected routes
	protected := router.Group("/api")
	protected.Use(middleware.JWTAuth())

	protected.POST("/categories", categoryHandler.CreateCategory)
	protected.GET("/categories", categoryHandler.GetAllCategory)
	protected.GET("/categories/:id", categoryHandler.GetCategoryById)
	protected.DELETE("/categories/:id", categoryHandler.DeleteCategory)
	protected.GET("/categories/:id/books", categoryHandler.GetBooksByCategory)
	protected.PUT("/categories/:id", categoryHandler.UpdateCategory)
}
