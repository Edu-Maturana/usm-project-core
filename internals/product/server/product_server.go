package server

import (
	"back-usm/internals/product/core/ports"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	productHandlers ports.ProductHandlers
}

func NewServer(productHandlers ports.ProductHandlers) *Server {
	return &Server{
		productHandlers: productHandlers,
	}
}

func (s *Server) Start() {
	app := fiber.New()
	api := app.Group("/api/v1")

	productRoutes := api.Group("/products")

	productRoutes.Get("/", s.productHandlers.GetAllProducts)
	productRoutes.Get("/:id", s.productHandlers.GetProduct)
	productRoutes.Post("/", s.productHandlers.CreateProduct)
	productRoutes.Put("/:id", s.productHandlers.UpdateProduct)
	productRoutes.Delete("/:id", s.productHandlers.DeleteProduct)

	err := app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server started on port 8080")
}