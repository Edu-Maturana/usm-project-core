package services

import (
	"back-usm/internals/comments/core/domain"
	"back-usm/internals/comments/core/ports"
)

type CommentServices struct {
	commentRepository ports.CommentRepository
}

func NewCommentServices(repository ports.CommentRepository) *CommentServices {
	return &CommentServices{
		commentRepository: repository,
	}
}

func (s *CommentServices) CreateComment(comment *domain.Comment) error {
	return s.commentRepository.Create(comment)
}

func (s *CommentServices) FindAllComments(productId string) ([]domain.Comment, error) {
	comments, err := s.commentRepository.FindAll(productId)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (s *CommentServices) DeleteComment(id string) error {
	return s.commentRepository.Delete(id)
}
