package ports

import (
	"back-usm/internals/product/core/domain"

	fiber "github.com/gofiber/fiber/v2"
)

type ProductServices interface {
	GetAllProducts() ([]domain.Product, error)
	GetProduct(id string) (domain.Product, error)
	CreateProduct(product domain.Product) (domain.Product, error)
	UpdateProduct(id string, product domain.Product) (domain.Product, error)
	DeleteProduct(id string) error
}

type ProductRepository interface {
	GetAll() ([]domain.Product, error)
	GetOne(id string) (domain.Product, error)
	Create(product domain.Product) (domain.Product, error)
	Update(id string, product domain.Product) (domain.Product, error)
	Delete(id string) error
}

type ProductHandlers struct {
	GetAllProducts (ctx *fiber.Ctx) error
	GetProduct     (ctx *fiber.Ctx) error
	CreateProduct  (ctx *fiber.Ctx) error
	UpdateProduct  (ctx *fiber.Ctx) error
	DeleteProduct  (ctx *fiber.Ctx) error
}