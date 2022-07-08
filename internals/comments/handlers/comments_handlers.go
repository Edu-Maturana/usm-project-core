package handlers

import (
	"back-usm/internals/comments/core/domain"
	"back-usm/internals/comments/core/ports"
	"back-usm/utils"

	"github.com/gofiber/fiber/v2"
)

type CommentHandlers struct {
	commentServices ports.CommentServices
}

func NewCommentHandlers(commentServices ports.CommentServices) *CommentHandlers {
	return &CommentHandlers{
		commentServices: commentServices,
	}
}

func (h *CommentHandlers) CreateComment(ctx *fiber.Ctx) error {
	comment := domain.Comment{}
	if err := ctx.BodyParser(&comment); err != nil {
		return err
	}

	validationError := utils.ValidateData(comment)
	if validationError != nil {
		return ctx.Status(400).JSON("Invalid data")
	}

	err := h.commentServices.CreateComment(&comment)
	if err != nil {
		return ctx.Status(404).JSON("Product not found")
	}

	return ctx.JSON(comment)
}

func (h *CommentHandlers) GetAllComments(ctx *fiber.Ctx) error {
	comments, err := h.commentServices.GetAllComments(ctx.Params("productId"))
	if err != nil {
		return err
	}

	return ctx.JSON(comments)
}

func (h *CommentHandlers) GetLastComments(ctx *fiber.Ctx) error {
	comments, err := h.commentServices.GetLastComments(ctx.Params("productId"))
	if err != nil {
		return err
	}

	return ctx.JSON(comments)
}

func (h *CommentHandlers) DeleteComment(ctx *fiber.Ctx) error {
	if err := h.commentServices.DeleteComment(ctx.Params("productId")); err != nil {
		return err
	}

	return ctx.JSON(map[string]string{"message": "Comment deleted"})
}
