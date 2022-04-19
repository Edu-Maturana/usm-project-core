package ports

import (
	"back-usm/internals/auth/core/domain"

	"github.com/gofiber/fiber/v2"
)

type AuthRepository interface {
	GetAll() ([]domain.Admin, error)
	GetOne(id string) (domain.Admin, error)
	Create(admin domain.Admin) (domain.Admin, error)
	Update(id string, admin domain.Admin) (domain.Admin, error)
	Delete(id string) error
}

type AuthServices interface {
	GetAllAdmins() ([]domain.Admin, error)
	GetOneAdmin(id string) (domain.Admin, error)
	CreateAdmin(admin domain.Admin) (domain.Admin, error)
	UpdateAdmin(id string, admin domain.Admin) (domain.Admin, error)
	DeleteAdmin(id string) error
	ActivateAccount(id string, admin domain.Admin) (domain.Admin, error)
	Login(admin domain.Admin) (domain.Admin, error)
}

type AuthHandlers interface {
	GetAllAdmins(c *fiber.Ctx) error
	GetOneAdmin(c *fiber.Ctx) error
	CreateAdmin(c *fiber.Ctx) error
	UpdateAdmin(c *fiber.Ctx) error
	DeleteAdmin(c *fiber.Ctx) error
	ActivateAccount(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
}
