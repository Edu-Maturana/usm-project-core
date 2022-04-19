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

func (s *AuthServices) GetAllAdmins() ([]domain.Admin, error) {
	users, err := s.authRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *AuthServices) GetOneAdmin(id string) (domain.Admin, error) {
	user, err := s.authRepository.GetOne(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *AuthServices) CreateAdmin(admin domain.Admin) (domain.Admin, error) {
	admin, err := s.authRepository.Create(admin)
	if err != nil {
		return admin, err
	}

	return admin, nil
}

func (s *AuthServices) UpdateAdmin(id string, admin domain.Admin) (domain.Admin, error) {
	admin, err := s.authRepository.Update(id, admin)
	if err != nil {
		return admin, err
	}

	return admin, nil
}

func (s *AuthServices) DeleteAdmin(id string) error {
	err := s.authRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthServices) ActivateAccount(id string, admin domain.Admin) (domain.Admin, error) {
	admin, err := s.authRepository.Update(id, admin)
	if err != nil {
		return admin, err
	}

	return admin, nil
}

func (s *AuthServices) Login(admin domain.Admin) (domain.Admin, error) {
	admin, err := s.authRepository.GetOne(admin.Email)
	if err != nil {
		return admin, err
	}

	return admin, nil
}
