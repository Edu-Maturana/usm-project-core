package repository

import (
	"back-usm/internals/auth/core/domain"
	"log"

	"github.com/fatih/color"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type AuthRepository struct {
	dsn string
	db  *gorm.DB
}

func NewAuthRepository(dsn string) *AuthRepository {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(domain.Admin{})

	log.Print(color.GreenString("Auth repository connected to database"))
	return &AuthRepository{
		dsn: dsn,
		db:  db,
	}
}

func (r *AuthRepository) GetAll() ([]domain.Admin, error) {
	var admins []domain.Admin
	err := r.db.Find(&admins).Error
	if err != nil {
		return admins, err
	}

	return admins, nil
}

func (r *AuthRepository) GetOne(id string) (domain.Admin, error) {
	var admin domain.Admin
	err := r.db.Where("id = ?", id).First(&admin).Error
	if err != nil {
		return admin, err
	}

	return admin, nil
}

func (r *AuthRepository) Create(admin domain.Admin) (domain.Admin, error) {
	err := r.db.Create(&admin).Error
	if err != nil {
		return admin, err
	}

	return admin, nil
}

func (r *AuthRepository) Update(id string, admin domain.Admin) (domain.Admin, error) {
	err := r.db.Model(&admin).Where("id = ?", id).Updates(&admin).Error
	if err != nil {
		return admin, err
	}

	return admin, nil
}

func (r *AuthRepository) Delete(id string) error {
	err := r.db.Where("id = ?", id).Delete(&domain.Admin{}).Error
	if err != nil {
		return err
	}

	return nil
}
