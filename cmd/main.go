package main

import (
	"back-usm/internals/product/core/services"
	"back-usm/internals/product/handlers"
	"back-usm/internals/product/repository"
	"back-usm/internals/product/server"
	"back-usm/utils"
)

func main() {
	dsn := utils.GetEnvVar("MYSQL_URI")

	productRepository := repository.NewProductRepository(dsn)
	productService := services.NewProductServices(productRepository)
	productHandlers := handlers.NewProductHandlers(productService)
	productServer := server.NewServer(productHandlers)
	productServer.Start()
}
