package ports

import (
	"back-usm/internals/order/core/domain"

	"github.com/gofiber/fiber/v2"
)

type OrderRepository interface {
	GetAll() ([]domain.Order, error)
	GetAllByCustomerEmail(customerEmail string) ([]domain.Order, error)
	GetOne(id string) (domain.Order, error)
	Create(order domain.Order) (domain.Order, error)
	Update(id string, order domain.Order) (domain.Order, error)
	Delete(id string) error
}

type OrderServices interface {
	GetAllOrders() ([]domain.Order, error)
	GetAllOrdersByCustomerEmail(customerEmail string) ([]domain.Order, error)
	GetOrder(id string) (domain.Order, error)
	CreateOrder(order domain.Order) (domain.Order, error)
	UpdateOrder(id string, order domain.Order) (domain.Order, error)
	DeleteOrder(id string) error
}

type OrderHandlers interface {
	GetAllOrders(ctx *fiber.Ctx) error
	GetAllOrdersByCustomerEmail(ctx *fiber.Ctx) error
	GetOrder(ctx *fiber.Ctx) error
	CreateOrder(ctx *fiber.Ctx) error
	UpdateOrder(ctx *fiber.Ctx) error
	DeleteOrder(ctx *fiber.Ctx) error
}
