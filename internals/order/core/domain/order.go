package domain

import (
	"gorm.io/gorm"
)

type Order struct {
	CustomerName    string      `gorm:"type:varchar(100)" json:"customer_name" validate:"required"`
	CustomerEmail   string      `gorm:"type:varchar(100)" json:"customer_email" validate:"required"`
	CustomerAddress string      `gorm:"type:varchar(100)" json:"customer_address" validate:"required"`
	OrderItems      []OrderItem `gorm:"foreignkey:OrderID" json:"order_items" validate:"required"`
	Status          string      `gorm:"type:varchar(100) default:pending" json:"status" enum:"pending,confirmed,paid,rejected"`
	Total           int32       `json:"total"`
	gorm.Model
}

type OrderItem struct {
	ProductID uint  `json:"product_id" validate:"required"`
	Quantity  int8  `json:"quantity" validate:"required"`
	Price     int32 `json:"price" validate:"required"`
	OrderID   uint  `json:"order_id" default:"0"`
}
