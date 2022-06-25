package services

import (
	"back-usm/internals/auth/core/domain"
	"back-usm/internals/auth/core/ports"
	"back-usm/utils"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
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

func (s *AuthServices) GetOneAdmin(email string) (domain.Admin, error) {
	user, err := s.authRepository.GetOne(email)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *AuthServices) CreateAdmin(newAdmin domain.Admin) (domain.Admin, error) {
	newAdmin.Password, _ = utils.EncryptPassword(newAdmin.Password)

	newAdmin, err := s.authRepository.Create(newAdmin)
	if err != nil {
		return newAdmin, err
	}

	return newAdmin, nil
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

func (s *AuthServices) GenerateToken(email string, id int) (string, error) {
	claims := jwt.MapClaims{
		"ID":    id,
		"email": email,
		"exp":   time.Now().AddDate(1, 0, 0), // 1 year of duration
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", fmt.Errorf("Error generating token")
	}

	return t, nil
}
