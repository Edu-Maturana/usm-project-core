package main

import (
	auth_services "back-usm/internals/auth/core/services"
	order_services "back-usm/internals/order/core/services"
	product_services "back-usm/internals/product/core/services"

	auth_handlers "back-usm/internals/auth/handlers"
	order_handlers "back-usm/internals/order/handlers"
	product_handlers "back-usm/internals/product/handlers"

	auth_repository "back-usm/internals/auth/repository"
	order_repository "back-usm/internals/order/repository"
	product_repository "back-usm/internals/product/repository"

	server "back-usm/cmd/server"
	"back-usm/utils"
)

func main() {
	dsn := utils.GetEnvVar("DSN")

	// Repositories
	authRepository := auth_repository.NewAuthRepository(dsn)
	orderRepository := order_repository.NewOrderRepository(dsn)
	productRepository := product_repository.NewProductRepository(dsn)

	// Services
	authServices := auth_services.NewAuthServices(authRepository)
	orderService := order_services.NewOrderServices(orderRepository)
	productService := product_services.NewProductServices(productRepository)

	// Handlers
	authHandlers := auth_handlers.NewAuthHandlers(authServices)
	orderHandlers := order_handlers.NewOrderHandlers(orderService)
	productHandlers := product_handlers.NewProductHandlers(productService)

	// Server
	server := server.NewServer(authHandlers, orderHandlers, productHandlers)

	// Init
	server.Start()
}
