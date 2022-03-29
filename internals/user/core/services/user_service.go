package services

import (
	"back-usm/internals/user/core/ports"
)

type ProductServices struct {
	productRepository ports.ProductRepository
}

func NewProductServices(repository ports.ProductRepository) *ProductServices {
	return &ProductServices{
		productRepository: repository,
	}
}

func (s *ProductServices) GetAllProducts() ([]struct{}, error) {
	products, err := s.productRepository.GetAllProducts()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *ProductServices) GetProduct(id string) (struct{}, error) {
	product, err := s.productRepository.GetProduct(id)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *ProductServices) CreateProduct(product struct{}) (struct{}, error) {
	product, err := s.productRepository.CreateProduct(product)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *ProductServices) UpdateProduct(id string, product struct{}) (struct{}, error) {
	product, err := s.productRepository.UpdateProduct(id, product)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *ProductServices) DeleteProduct(id string) error {
	err := s.productRepository.DeleteProduct(id)
	if err != nil {
		return err
	}

	return nil
}
