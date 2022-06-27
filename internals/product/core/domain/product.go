package domain

import "gorm.io/gorm"

type Product struct {
	Name        string `gorm:"type:varchar(100)" json:"name" validate:"required"`
	Description string `gorm:"type:varchar(255)" json:"description" validate:"required"`
	Image       string `gorm:"type:varchar(100)" json:"image" validate:"required"`
	Stock       uint8  `gorm:"type:tinyint(2)" json:"stock" validate:"required,numeric,min=0,max=50"`
	Price       uint16 `gorm:"type:smallint(2)" json:"price" validate:"required,numeric,min=1000,max=50000"`
	gorm.Model
}
