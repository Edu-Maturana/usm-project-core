package handlers

import (
	"back-usm/internals/product/core/domain"
	"back-usm/internals/product/core/ports"

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
		return err
	}

	return ctx.JSON(products)
}

func (h *ProductHandlers) GetProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	product, err := h.productServices.GetProduct(id)
	if err != nil {
		return err
	}

	return ctx.JSON(product)
}

func (h *ProductHandlers) CreateProduct(ctx *fiber.Ctx) error {
	var product domain.Product
	if err := ctx.BodyParser(&product); err != nil {
		return err
	}

	product, err := h.productServices.CreateProduct(product)
	if err != nil {
		return err
	}

	return ctx.JSON(product)
}

func (h *ProductHandlers) UpdateProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var product domain.Product
	if err := ctx.BodyParser(&product); err != nil {
		return err
	}

	product, err := h.productServices.UpdateProduct(id, product)
	if err != nil {
		return err
	}

	return ctx.JSON(product)
}

func (h *ProductHandlers) DeleteProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := h.productServices.DeleteProduct(id)
	if err != nil {
		return err
	}

	return nil
}
