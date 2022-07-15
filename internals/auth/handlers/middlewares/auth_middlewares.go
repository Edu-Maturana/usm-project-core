package middlewares

import (
	"back-usm/internals/auth/core/domain"
	"back-usm/internals/auth/core/ports"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type AuthMiddlewares struct {
	authServices ports.AuthServices
}

func NewAuthHandlers(authServices ports.AuthServices) *AuthMiddlewares {
	return &AuthMiddlewares{
		authServices: authServices,
	}
}

func (m *AuthMiddlewares) VerifyIfAdminExists(ctx *fiber.Ctx) error {
	email := ctx.Params("email")
	_, err := m.authServices.GetOneAdmin(email)
	if err != nil {
		return ctx.Status(404).JSON("Admin not found")
	}

	return ctx.Next()
}

func (m *AuthMiddlewares) VerifyIfAdminIsNew(ctx *fiber.Ctx) error {
	var admin domain.Admin
	ctx.BodyParser(&admin)
	_, err := m.authServices.GetOneAdmin(admin.Email)
	if err == nil {
		return ctx.Status(400).JSON("Admin already exists")
	}

	return ctx.Next()
}

func (m *AuthMiddlewares) ValidateToken(ctx *fiber.Ctx) error {
	token := ctx.Get("token")

	if token == "" {
		return ctx.Status(400).JSON("There is not token")
	}

	bytes, err := jwt.Parse(token, nil)
	if err != nil {
		return ctx.Status(400).JSON("Invalid token")
	}

	claims := bytes.Claims.(jwt.MapClaims)

	_, err = m.authServices.GetOneAdmin(claims["email"].(string))
	if err != nil {
		return ctx.Status(400).JSON("Unauthorized")
	}

	return ctx.Next()
}
