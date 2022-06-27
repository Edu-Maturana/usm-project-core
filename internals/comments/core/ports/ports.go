package ports

import (
	"back-usm/internals/comments/core/domain"

	"github.com/gofiber/fiber/v2"
)

type CommentRepository interface {
	Create(comment *domain.Comment) error
	FindAll(productId string) ([]domain.Comment, error)
	Delete(id string) error
}

type CommentServices interface {
	CreateComment(comment *domain.Comment) error
	FindAllComments(productId string) ([]domain.Comment, error)
	DeleteComment(productId string) error
}

type CommentHandlers interface {
	CreateComment(ctx *fiber.Ctx) error
	FindAllComments(ctx *fiber.Ctx) error
	DeleteComment(ctx *fiber.Ctx) error
}
