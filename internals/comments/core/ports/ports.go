package ports

import (
	"back-usm/internals/comments/core/domain"

	"github.com/gofiber/fiber/v2"
)

type CommentRepository interface {
	Create(comment *domain.Comment) error
	GetAll(productId string) ([]domain.Comment, error)
	GetLast(productId string) ([]domain.Comment, error)
	Delete(id string) error
}

type CommentServices interface {
	CreateComment(comment *domain.Comment) error
	GetAllComments(productId string) ([]domain.Comment, error)
	GetLastComments(productId string) ([]domain.Comment, error)
	DeleteComment(productId string) error
}

type CommentHandlers interface {
	CreateComment(ctx *fiber.Ctx) error
	GetAllComments(ctx *fiber.Ctx) error
	GetLastComments(ctx *fiber.Ctx) error
	DeleteComment(ctx *fiber.Ctx) error
}
