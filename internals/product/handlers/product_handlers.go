package handlers

import (
	"back-usm/internals/product/core/domain"
	"back-usm/internals/product/core/ports"
	"back-usm/utils"

	"github.com/gofiber/fiber/v2"
)

type ProductHandlers struct {
	productServices ports.ProductServices
}

func NewProductHandlers(productServices ports.ProductServices) *ProductHandlers {
	return &ProductHandlers{
		productServices: productServices,
	}
}

func (h *ProductHandlers) GetAllProducts(ctx *fiber.Ctx) error {
	products, err := h.productServices.GetAllProducts()
	if err != nil {
		return ctx.Status(404).JSON("Products not found")
	}

	return ctx.JSON(products)
}

func (h *ProductHandlers) GetProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	product, err := h.productServices.GetProduct(id)
	if err != nil {
		return ctx.Status(404).JSON("Product not found")
	}

	return ctx.Status(200).JSON(product)
}

func (h *ProductHandlers) CreateProduct(ctx *fiber.Ctx) error {
	var product domain.Product
	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(400).JSON("Invalid product")
	}

	validationError := utils.ValidateData(product)
	if validationError != nil {
		return ctx.Status(400).JSON("Invalid data")
	}

	product, err := h.productServices.CreateProduct(product)
	if err != nil {
		return ctx.Status(400).JSON(err)
	}

	return ctx.Status(201).JSON(product)
}

func (h *ProductHandlers) UpdateProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var product domain.Product
	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(400).JSON("Invalid data")
	}

	product, err := h.productServices.UpdateProduct(id, product)
	if err != nil {
		return ctx.Status(400).JSON("Error updating product")
	}

	return ctx.Status(200).JSON("Product updated successfully")
}

func (h *ProductHandlers) DeleteProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := h.productServices.DeleteProduct(id)
	if err != nil {
		return ctx.Status(400).JSON("Error deleting product")
	}

	return ctx.Status(200).JSON("Product deleted successfully")
}
