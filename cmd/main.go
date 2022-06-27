package main

import (
	auth_services "back-usm/internals/auth/core/services"
	comments_services "back-usm/internals/comments/core/services"
	product_services "back-usm/internals/product/core/services"

	auth_handlers "back-usm/internals/auth/handlers"
	auth_middlewares "back-usm/internals/auth/handlers/middlewares"
	comments_handlers "back-usm/internals/comments/handlers"
	product_handlers "back-usm/internals/product/handlers"

	auth_repository "back-usm/internals/auth/repository"
	comments_repository "back-usm/internals/comments/repository"
	product_repository "back-usm/internals/product/repository"

	server "back-usm/cmd/server"
	"back-usm/utils"
)

func main() {
	dsn := utils.GetEnvVar("DSN")

	// Repositories
	authRepository := auth_repository.NewAuthRepository(dsn)
	productRepository := product_repository.NewProductRepository(dsn)
	commentsRepository := comments_repository.NewCommentRepository(dsn)

	// Services
	authServices := auth_services.NewAuthServices(authRepository)
	productService := product_services.NewProductServices(productRepository)
	commentsService := comments_services.NewCommentServices(commentsRepository, productRepository)

	// Middlewares
	authMiddlewares := auth_middlewares.NewAuthHandlers(authServices)

	// Handlers
	authHandlers := auth_handlers.NewAuthHandlers(authServices)
	productHandlers := product_handlers.NewProductHandlers(productService)
	commentsHandlers := comments_handlers.NewCommentHandlers(commentsService)

	// Server
	server := server.NewServer(authHandlers, productHandlers, commentsHandlers, authMiddlewares)

	// Init
	server.Start()
}
