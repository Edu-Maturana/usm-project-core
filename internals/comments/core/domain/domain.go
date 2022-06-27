package domain

import "gorm.io/gorm"

type Comment struct {
	Customer  string `json:"customer" validate:"required"`
	Content   string `json:"content" validate:"required"`
	Stars     int    `json:"stars" validate:"required,numeric,min=1,max=5"`
	ProductId uint   `json:"product_id" validate:"required,numeric"`
	gorm.Model
}
