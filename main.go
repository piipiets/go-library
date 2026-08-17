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

	migration.Initiator(connection.DBConnections)

	// repository
	userRepository := repository.NewUserRepository(connection.DBConnections)
	categoryRespository := repository.NewCategoryRepository(connection.DBConnections)

	// service
	authService := service.NewAuthService(userRepository)
	categoryService := service.NewCategoryService(categoryRespository)
	// handler
	authHandler := handler.NewAuthHandler(authService)
	categoryHandler := handler.NewCategoryHancder(categoryService)

	// router
	router := gin.Default()

	// routes
	routes.SetupRoutes(
		router,
		authHandler,
		categoryHandler,
	)

	router.Run(":8080")
}
