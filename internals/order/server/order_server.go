package server

import (
	"back-usm/internals/order/core/ports"
	"log"

	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	orderHandlers ports.OrderHandlers
}

func NewServer(orderHandlers ports.OrderHandlers) *Server {
	return &Server{
		orderHandlers: orderHandlers,
	}
}

func (s *Server) Start() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	api := app.Group("/api/v1")

	orderRoutes := api.Group("/orders")

	orderRoutes.Get("/", s.orderHandlers.GetAllOrders)
	orderRoutes.Get("/:id", s.orderHandlers.GetOrder)
	orderRoutes.Post("/", s.orderHandlers.CreateOrder)
	orderRoutes.Put("/:id", s.orderHandlers.UpdateOrder)
	orderRoutes.Delete("/:id", s.orderHandlers.DeleteOrder)

	log.Println(color.HiBlueString("Server listening on port 8081"))
	app.Listen(":8081")
}
