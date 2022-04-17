package services

import (
	"back-usm/internals/order/core/domain"
	"back-usm/internals/order/core/ports"
)

type OrderServices struct {
	orderRepository ports.OrderRepository
}

func NewOrderServices(repository ports.OrderRepository) *OrderServices {
	return &OrderServices{
		orderRepository: repository,
	}
}

func (s *OrderServices) GetAllOrders() ([]domain.Order, error) {
	orders, err := s.orderRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *OrderServices) GetOrder(id string) (domain.Order, error) {
	order, err := s.orderRepository.GetOne(id)
	if err != nil {
		return domain.Order{}, err
	}

	return order, nil
}

func (s *OrderServices) CreateOrder(order domain.Order) (domain.Order, error) {
	for _, item := range order.OrderItems {
		order.Total += int32(item.Price) * int32(item.Quantity)
	}

	order, err := s.orderRepository.Create(order)
	if err != nil {
		return domain.Order{}, err
	}

	return order, nil
}

func (s *OrderServices) UpdateOrder(id string, order domain.Order) (domain.Order, error) {
	order, err := s.orderRepository.Update(id, order)
	if err != nil {
		return domain.Order{}, err
	}

	return order, nil
}

func (s *OrderServices) DeleteOrder(id string) error {
	err := s.orderRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
