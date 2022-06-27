package services

import (
	"back-usm/internals/comments/core/domain"
	commentsPorts "back-usm/internals/comments/core/ports"
	productsPorts "back-usm/internals/product/core/ports"
	"fmt"
)

type CommentServices struct {
	commentRepository commentsPorts.CommentRepository
	productRepository productsPorts.ProductRepository
}

func NewCommentServices(repository commentsPorts.CommentRepository, productRepository productsPorts.ProductRepository) *CommentServices {
	return &CommentServices{
		commentRepository: repository,
		productRepository: productRepository,
	}
}

func (s *CommentServices) CreateComment(comment *domain.Comment) error {
	product, _ := s.productRepository.GetOne(fmt.Sprint(comment.ProductId))
	if product.ID == 0 {
		return fmt.Errorf("Product not found")
	}

	err := s.commentRepository.Create(comment)
	if err != nil {
		return err
	}

	return nil
}

func (s *CommentServices) FindAllComments(productId string) ([]domain.Comment, error) {
	comments, err := s.commentRepository.FindAll(productId)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (s *CommentServices) DeleteComment(id string) error {
	err := s.commentRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
