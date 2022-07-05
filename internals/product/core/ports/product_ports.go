package ports

import (
	"back-usm/internals/product/core/domain"

	"github.com/gofiber/fiber/v2"
)

type ProductRepository interface {
	GetAll(priceSort int) ([]domain.Product, error)
	GetOne(id string) (domain.Product, error)
	Create(product domain.Product) (domain.Product, error)
	Update(id string, product domain.Product) (domain.Product, error)
	Delete(id string) error
}

type ProductServices interface {
	GetAllProducts(priceSort int) ([]domain.Product, error)
	GetProduct(id string) (domain.Product, error)
	CreateProduct(product domain.Product) (domain.Product, error)
	UpdateProduct(id string, product domain.Product) (domain.Product, error)
	DeleteProduct(id string) error
}

type ProductHandlers interface {
	GetAllProducts(ctx *fiber.Ctx) error
	GetProduct(ctx *fiber.Ctx) error
	CreateProduct(ctx *fiber.Ctx) error
	UpdateProduct(ctx *fiber.Ctx) error
	DeleteProduct(ctx *fiber.Ctx) error
}
