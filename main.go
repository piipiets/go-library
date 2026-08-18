package main

import (
	"github.com/gin-gonic/gin"

	"github.com/piipiets/go-library/config"
	"github.com/piipiets/go-library/database/connection"
	"github.com/piipiets/go-library/database/migration"
	"github.com/piipiets/go-library/handler"
	"github.com/piipiets/go-library/repository"
	"github.com/piipiets/go-library/routes"
	"github.com/piipiets/go-library/service"
)

func main() {
	config.Initiator()

	connection.Initiator()
	defer connection.DBConnections.Close()

	conn := connection.DBConnections
	migration.Initiator(conn)

	// repository
	userRepository := repository.NewUserRepository(conn)
	categoryRespository := repository.NewCategoryRepository(conn)
	bookRepository := repository.NewBookRepository(conn)

	// service
	authService := service.NewAuthService(userRepository)
	categoryService := service.NewCategoryService(categoryRespository)
	userService := service.NewUserService(userRepository)
	bookService := service.NewBookService(bookRepository)

	// handler
	authHandler := handler.NewAuthHandler(authService)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	userHandler := handler.NewUserHandler(userService)
	bookHandler := handler.NewBookHandler(bookService)

	// router
	router := gin.Default()

	// routes
	routes.SetupRoutes(
		router,
		authHandler,
		categoryHandler,
		userHandler,
		bookHandler,
	)

	router.Run(":8080")
}
