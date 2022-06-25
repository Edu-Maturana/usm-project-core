package handlers

import (
	"back-usm/internals/auth/core/domain"
	"back-usm/internals/auth/core/ports"
	"back-usm/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthHandlers struct {
	authServices ports.AuthServices
}

func NewAuthHandlers(authServices ports.AuthServices) *AuthHandlers {
	return &AuthHandlers{
		authServices: authServices,
	}
}

func (h *AuthHandlers) GetAllAdmins(c *fiber.Ctx) error {
	users, err := h.authServices.GetAllAdmins()
	if err != nil {
		return c.Status(404).JSON("Admins not found")
	}

	return c.Status(200).JSON(users)
}

func (h *AuthHandlers) GetOneAdmin(c *fiber.Ctx) error {
	email := c.Params("email")
	user, err := h.authServices.GetOneAdmin(email)
	if err != nil {
		return c.Status(404).JSON("Admin not found")
	}

	return c.Status(200).JSON(user)
}

func (h *AuthHandlers) CreateAdmin(c *fiber.Ctx) error {
	var admin domain.Admin
	if err := c.BodyParser(&admin); err != nil {
		return c.Status(400).JSON("Invalid admin")
	}

	if err := utils.ValidateData(admin); err != nil {
		return c.Status(400).JSON("Invalid data")
	}

	admin, err := h.authServices.CreateAdmin(admin)
	if err != nil {
		return c.Status(400).JSON(err)
	}

	return c.Status(201).JSON(admin)
}

func (h *AuthHandlers) UpdateAdmin(c *fiber.Ctx) error {
	var admin domain.Admin
	email := c.Params("email")

	if err := c.BodyParser(&admin); err != nil {
		return c.Status(400).JSON("Invalid admin")
	}

	admin, err := h.authServices.UpdateAdmin(email, admin)
	if err != nil {
		return c.Status(400).JSON("Error updating admin")
	}

	return c.Status(200).JSON("Admin updated")
}

func (h *AuthHandlers) DeleteAdmin(c *fiber.Ctx) error {
	email := c.Params("email")
	err := h.authServices.DeleteAdmin(email)
	if err != nil {
		return c.Status(400).JSON("Error deleting admin")
	}

	return c.Status(200).JSON("Admin deleted")
}

func (h *AuthHandlers) ActivateAccount(c *fiber.Ctx) error {
	var admin domain.Admin
	email := c.Params("email")

	if err := c.BodyParser(&admin); err != nil {
		return c.Status(400).JSON("Invalid admin")
	}

	admin, err := h.authServices.ActivateAccount(email, admin)
	if err != nil {
		return c.Status(400).JSON("Error activating admin")
	}

	return c.Status(200).JSON("Account activated")
}

func (h *AuthHandlers) Login(c *fiber.Ctx) error {
	var admin domain.Admin
	if err := c.BodyParser(&admin); err != nil {
		return c.Status(400).JSON("Invalid credentials")
	}

	admin, err := h.authServices.Login(admin)
	if err != nil {
		return c.Status(400).JSON("Invalid credentials")
	}

	token, err := h.authServices.GenerateToken(admin.Email, int(admin.ID))
	if err != nil {
		return c.Status(400).JSON(err)
	}

	return c.JSON(map[string]string{
		"message": "Login succesful",
		"token":   token,
	})
}
