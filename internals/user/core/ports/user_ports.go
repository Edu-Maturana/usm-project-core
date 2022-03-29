package ports

import (
	"back-usm/internals/core/user/domain"

	fiber "github.com/gofiber/fiber/v2"
)

type ProductServices interface {
	GetAllProducts() ([]struct{}, error)
	GetProduct(id string) (struct{}, error)
	CreateProduct(product struct{}) (struct{}, error)
	UpdateProduct(id string, product struct{}) (struct{}, error)
	DeleteProduct(id string) error
}

type ProductRepository interface {
	GetAllProducts() ([]struct{}, error)
	GetProduct(id string) (struct{}, error)
	CreateProduct(product struct{}) (struct{}, error)
	UpdateProduct(id string, product struct{}) (struct{}, error)
	DeleteProduct(id string) error
}

type ProductHandlers struct {
	GetAllProducts (ctx *fiber.Ctx) error
	GetProduct     (ctx *fiber.Ctx) error
	CreateProduct  (ctx *fiber.Ctx) error
	UpdateProduct  (ctx *fiber.Ctx) error
	DeleteProduct  (ctx *fiber.Ctx) error
}