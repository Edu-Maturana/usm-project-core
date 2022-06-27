package server

import (
	auth "back-usm/internals/auth/core/ports"
	comments "back-usm/internals/comments/core/ports"
	products "back-usm/internals/product/core/ports"
	"log"

	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	authHandlers     auth.AuthHandlers
	productHandlers  products.ProductHandlers
	commentsHandlers comments.CommentHandlers
	authMiddlewares  auth.AuthMiddlewares
}

func NewServer(
	auth auth.AuthHandlers,
	products products.ProductHandlers,
	comments comments.CommentHandlers,
	authMiddlewares auth.AuthMiddlewares) *Server {
	return &Server{
		authHandlers:     auth,
		productHandlers:  products,
		commentsHandlers: comments,
		authMiddlewares:  authMiddlewares,
	}
}

func (s *Server) Start() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Use(logger.New(logger.Config{
		Format: "${time} ${status} ${method} ${path} ${latency}\n",
	}))

	app.Use(cors.New())

	api := app.Group("/api/v1")

	authRoutes := api.Group("/auth")
	productRoutes := api.Group("/products")
	commentRoutes := api.Group("/comments")

	authRoutes.Get("/admins", s.authMiddlewares.ValidateToken, s.authHandlers.GetAllAdmins)
	authRoutes.Get("/admins/:email", s.authHandlers.GetOneAdmin)
	authRoutes.Post("/admins", s.authMiddlewares.VerifyIfAdminIsNew, s.authHandlers.CreateAdmin)
	authRoutes.Put("/admins/:email", s.authHandlers.UpdateAdmin)
	authRoutes.Delete("/admins/:email", s.authHandlers.DeleteAdmin)

	authRoutes.Put("/activate/:id", s.authHandlers.ActivateAccount)
	authRoutes.Post("/login", s.authHandlers.Login)

	productRoutes.Get("/", s.productHandlers.GetAllProducts)
	productRoutes.Get("/:id", s.productHandlers.GetProduct)
	productRoutes.Post("/", s.authMiddlewares.ValidateToken, s.productHandlers.CreateProduct)
	productRoutes.Put("/:id", s.productHandlers.UpdateProduct)
	productRoutes.Delete("/:id", s.productHandlers.DeleteProduct)

	commentRoutes.Get("/:productId", s.commentsHandlers.FindAllComments)
	commentRoutes.Post("/", s.commentsHandlers.CreateComment)
	commentRoutes.Delete("/:productId", s.commentsHandlers.DeleteComment)

	log.Println(color.BlueString("Server listening on port 8080"))
	app.Listen(":8080")
}
