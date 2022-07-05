package repository

import (
	"back-usm/internals/product/core/domain"
	"log"

	"github.com/fatih/color"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ProductRepository struct {
	dsn string
	db  *gorm.DB
}

func NewProductRepository(dsn string) *ProductRepository {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(domain.Product{})

	log.Print(color.GreenString("Products repository connected to database"))
	return &ProductRepository{
		dsn: dsn,
		db:  db,
	}
}

func (r *ProductRepository) GetAll(priceSort int) ([]domain.Product, error) {
	var products []domain.Product

	if priceSort > 0 {
		if priceSort == 1 {
			err := r.db.Order("price ASC").Find(&products).Error
			if err != nil {
				return nil, err
			}
		} else if priceSort == 2 {
			err := r.db.Order("price DESC").Find(&products).Error
			if err != nil {
				return nil, err
			}
		}

		return products, nil
	}

	err := r.db.Order("created_at desc").Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetOne(id string) (domain.Product, error) {
	var product domain.Product
	err := r.db.Where("id = ?", id).First(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *ProductRepository) Create(product domain.Product) (domain.Product, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *ProductRepository) Update(id string, product domain.Product) (domain.Product, error) {
	err := r.db.Model(&product).Where("id = ?", id).Updates(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *ProductRepository) Delete(id string) error {
	err := r.db.Where("id = ?", id).Delete(&domain.Product{}).Error
	if err != nil {
		return err
	}

	return nil
}
