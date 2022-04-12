package handlers

import (
	"back-usm/internals/order/core/domain"
	"back-usm/internals/order/core/ports"
	"back-usm/utils"

	"github.com/gofiber/fiber/v2"
)

type OrderHandlers struct {
	orderServices ports.OrderServices
}

func NewOrderHandlers(orderServices ports.OrderServices) *OrderHandlers {
	return &OrderHandlers{
		orderServices: orderServices,
	}
}

func (h *OrderHandlers) GetAllOrders(ctx *fiber.Ctx) error {
	orders, err := h.orderServices.GetAllOrders()
	if err != nil {
		return ctx.Status(404).JSON("Orders not found")
	}
	return ctx.JSON(orders)
}

func (h *OrderHandlers) GetOrder(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	order, err := h.orderServices.GetOrder(id)
	if err != nil {
		return ctx.Status(404).JSON("Order not found")
	}

	return ctx.JSON(order)
}

func (h *OrderHandlers) CreateOrder(ctx *fiber.Ctx) error {
	var order domain.Order
	if err := ctx.BodyParser(&order); err != nil {
		return ctx.Status(400).JSON(err)
	}

	validationError := utils.ValidateData(order)
	if validationError != nil {
		return ctx.Status(400).JSON("Invalid data, all fields are required")
	}
	order, err := h.orderServices.CreateOrder(order)
	if err != nil {
		return ctx.Status(400).JSON(err)
	}

	return ctx.JSON(order)
}

func (h *OrderHandlers) UpdateOrder(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var order domain.Order
	if err := ctx.BodyParser(&order); err != nil {
		return ctx.Status(400).JSON("Invalid order")
	}

	validationError := utils.ValidateData(order)
	if validationError != nil {
		return ctx.Status(400).JSON("Invalid data, all fields are required")
	}

	order, err := h.orderServices.UpdateOrder(id, order)
	if err != nil {
		return ctx.Status(400).JSON("Invalid order")
	}

	return ctx.JSON(order)
}

func (h *OrderHandlers) DeleteOrder(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := h.orderServices.DeleteOrder(id)
	if err != nil {
		return ctx.Status(404).JSON("Order not found")
	}

	return ctx.JSON("Order deleted")
}
