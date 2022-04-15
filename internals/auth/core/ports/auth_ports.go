package ports

import (
	"back-usm/internals/auth/core/domain"

	"github.com/gofiber/fiber/v2"
)

type AuthRepository interface {
	Login(admin domain.Admin) (domain.Admin, error)
	ActivateAccount(admin domain.Admin) (domain.Admin, error)
}

type AuthServices interface {
	Login(admin domain.Admin) (domain.Admin, error)
	ActivateAccount(admin domain.Admin) (domain.Admin, error)
}

type AuthHandlers interface {
	Login(ctx *fiber.Ctx) error
	ActivateAccount(admin domain.Admin) (domain.Admin, error)
}
