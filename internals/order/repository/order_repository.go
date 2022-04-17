package repository

import (
	"back-usm/internals/order/core/domain"
	"log"

	"github.com/fatih/color"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type OrderRepository struct {
	dsn string
	db  *gorm.DB
}

func NewOrderRepository(dsn string) *OrderRepository {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(domain.Order{})
	db.AutoMigrate(domain.OrderItem{})

	log.Print(color.GreenString("Orders repository connected to database"))
	return &OrderRepository{
		dsn: dsn,
		db:  db,
	}
}

func (r *OrderRepository) GetAll() ([]domain.Order, error) {
	var orders []domain.Order
	err := r.db.Find(&orders).Error
	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (r *OrderRepository) GetOne(id string) (domain.Order, error) {
	var order domain.Order
	err := r.db.Where("id = ?", id).First(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *OrderRepository) Create(order domain.Order) (domain.Order, error) {
	err := r.db.Create(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *OrderRepository) Update(id string, order domain.Order) (domain.Order, error) {
	err := r.db.Model(&order).Where("id = ?", id).Updates(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *OrderRepository) Delete(id string) error {
	err := r.db.Where("id = ?", id).Delete(&domain.Order{}).Error
	if err != nil {
		return err
	}

	return nil
}
