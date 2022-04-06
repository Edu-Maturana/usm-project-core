package server

import (
	order_ports "back-usm/internals/order/core/ports"
	products_port "back-usm/internals/product/core/ports"
	"log"

	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	productHandlers products_port.ProductHandlers
	orderHandlers   order_ports.OrderHandlers
}

func NewServer(productHandlers products_port.ProductHandlers, orderHandlers order_ports.OrderHandlers) *Server {
	return &Server{
		productHandlers: productHandlers,
		orderHandlers:   orderHandlers,
	}
}

func (s *Server) Start() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Use(logger.New(logger.Config{
		Format: "${time} ${status} ${method} ${path} ${latency}\n",
	}))

	api := app.Group("/api/v1")

	productRoutes := api.Group("/products")
	orderRoutes := api.Group("/orders")

	productRoutes.Get("/", s.productHandlers.GetAllProducts)
	productRoutes.Get("/:id", s.productHandlers.GetProduct)
	productRoutes.Post("/", s.productHandlers.CreateProduct)
	productRoutes.Put("/:id", s.productHandlers.UpdateProduct)
	productRoutes.Delete("/:id", s.productHandlers.DeleteProduct)

	orderRoutes.Get("/", s.orderHandlers.GetAllOrders)
	orderRoutes.Get("/:id", s.orderHandlers.GetOrder)
	orderRoutes.Post("/", s.orderHandlers.CreateOrder)
	orderRoutes.Put("/:id", s.orderHandlers.UpdateOrder)
	orderRoutes.Delete("/:id", s.orderHandlers.DeleteOrder)

	log.Println(color.HiBlueString("Server listening on port 8080"))
	app.Listen(":8080")
}
