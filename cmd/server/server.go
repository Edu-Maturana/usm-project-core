package server

import (
	auth "back-usm/internals/auth/core/ports"
	order "back-usm/internals/order/core/ports"
	products "back-usm/internals/product/core/ports"
	"log"

	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	authHandlers    auth.AuthHandlers
	orderHandlers   order.OrderHandlers
	productHandlers products.ProductHandlers
	authMiddlewares auth.AuthMiddlewares
}

func NewServer(auth auth.AuthHandlers, orders order.OrderHandlers, products products.ProductHandlers, authMiddlewares auth.AuthMiddlewares) *Server {
	return &Server{
		authHandlers:    auth,
		orderHandlers:   orders,
		productHandlers: products,
		authMiddlewares: authMiddlewares,
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

	authRoutes := api.Group("/auth")
	orderRoutes := api.Group("/orders")
	productRoutes := api.Group("/products")

	authRoutes.Get("/admins", s.authHandlers.GetAllAdmins)
	authRoutes.Get("/admins/:email", s.authHandlers.GetOneAdmin)
	authRoutes.Post("/admins", s.authMiddlewares.VerifyIfAdminIsNew, s.authHandlers.CreateAdmin)
	authRoutes.Put("/admins/:email", s.authHandlers.UpdateAdmin)
	authRoutes.Delete("/admins/:email", s.authHandlers.DeleteAdmin)

	authRoutes.Put("/activate/:id", s.authHandlers.ActivateAccount)
	authRoutes.Post("/login", s.authHandlers.Login)

	orderRoutes.Get("/", s.orderHandlers.GetAllOrders)
	orderRoutes.Get("/:id", s.orderHandlers.GetOrder)
	orderRoutes.Post("/", s.orderHandlers.CreateOrder)
	orderRoutes.Put("/:id", s.orderHandlers.UpdateOrder)
	orderRoutes.Delete("/:id", s.orderHandlers.DeleteOrder)

	productRoutes.Get("/", s.productHandlers.GetAllProducts)
	productRoutes.Get("/:id", s.productHandlers.GetProduct)
	productRoutes.Post("/", s.productHandlers.CreateProduct)
	productRoutes.Put("/:id", s.productHandlers.UpdateProduct)
	productRoutes.Delete("/:id", s.productHandlers.DeleteProduct)

	log.Println(color.BlueString("Server listening on port 8080"))
	app.Listen(":8080")
}
