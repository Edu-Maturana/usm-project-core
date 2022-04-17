package services

import (
	"back-usm/internals/auth/core/domain"
	"back-usm/internals/auth/core/ports"
)

type AuthServices struct {
	authRepository ports.AuthRepository
}

func NewAuthServices(repository ports.AuthRepository) *AuthServices {
	return &AuthServices{
		authRepository: repository,
	}
}

func (s *AuthServices) GetAllUsers() ([]domain.Admin, error) {
	users, err := s.authRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *AuthServices) GetOneUser(id string) (domain.Admin, error) {
	user, err := s.authRepository.GetOne(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *AuthServices) CreateUser(user domain.Admin) (domain.Admin, error) {
	user, err := s.authRepository.Create(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *AuthServices) UpdateUser(user domain.Admin) (domain.Admin, error) {
	user, err := s.authRepository.Update(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *AuthServices) DeleteUser(id string) error {
	err := s.authRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthServices) Login(user domain.Admin) (domain.Admin, error) {
	user, err := s.authRepository.Login(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *AuthServices) ActivateAccount(user domain.Admin) (domain.Admin, error) {
	user, err := s.authRepository.ActivateAccount(user)
	if err != nil {
		return user, err
	}

	return user, nil
}
