package ports

import (
	"back-usm/internals/core/user/domain"

	fiber "github.com/gofiber/fiber/v2"
)

type ProductService interface {
	GetAllProducts() ([]domain.Product, error)
	GetProduct(id string) (domain.Product, error)
	CreateProduct(product domain.Product) (domain.Product, error)
	UpdateProduct(id string, product domain.Product) (domain.Product, error)
	DeleteProduct(id string) error
}

type ProductRepository interface {
	GetAllProducts() ([]domain.Product, error)
	GetProduct(id string) (domain.Product, error)
	CreateProduct(product domain.Product) (domain.Product, error)
	UpdateProduct(id string, product domain.Product) (domain.Product, error)
	DeleteProduct(id string) error
}

type ProductHandlers struct {
	GetAllProducts (ctx *fiber.Ctx) error
	GetProduct     (ctx *fiber.Ctx) error
	CreateProduct  (ctx *fiber.Ctx) error
	UpdateProduct  (ctx *fiber.Ctx) error
	DeleteProduct  (ctx *fiber.Ctx) error
}