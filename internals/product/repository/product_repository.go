package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ProductRepository struct {
	uri   string
	db    *gorm.DB
	table string
}

func NewProductRepository(uri string) (*ProductRepository, error) {
	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &ProductRepository{
		uri:   uri,
		db:    db,
		table: "products",
	}, nil
}

func (r *ProductRepository) GetAll() ([]struct{}, error) {
	var products []struct{}
	err := r.db.Table(r.table).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetOne(id string) (struct{}, error) {
	var product struct{}
	err := r.db.Table(r.table).Where("id = ?", id).First(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *ProductRepository) Create(product struct{}) (struct{}, error) {
	err := r.db.Table(r.table).Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *ProductRepository) Update(id string, product struct{}) (struct{}, error) {
	err := r.db.Table(r.table).Where("id = ?", id).Updates(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *ProductRepository) Delete(id string) error {
	err := r.db.Table(r.table).Where("id = ?", id).Delete(&struct{}{}).Error
	if err != nil {
		return err
	}

	return nil
}
