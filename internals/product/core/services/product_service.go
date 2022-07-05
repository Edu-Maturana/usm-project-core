package services

import (
	"back-usm/internals/product/core/domain"
	"back-usm/internals/product/core/ports"
)

type ProductServices struct {
	productRepository ports.ProductRepository
}

func NewProductServices(repository ports.ProductRepository) *ProductServices {
	return &ProductServices{
		productRepository: repository,
	}
}

func (s *ProductServices) GetAllProducts(priceSort int) ([]domain.Product, error) {
	products, err := s.productRepository.GetAll(priceSort)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *ProductServices) GetProduct(id string) (domain.Product, error) {
	product, err := s.productRepository.GetOne(id)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *ProductServices) CreateProduct(product domain.Product) (domain.Product, error) {
	product, err := s.productRepository.Create(product)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *ProductServices) UpdateProduct(id string, product domain.Product) (domain.Product, error) {
	product, err := s.productRepository.Update(id, product)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *ProductServices) DeleteProduct(id string) error {
	err := s.productRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
