package domain

import "gorm.io/gorm"

type Comment struct {
	Customer  string `json:"customer"`
	Content   string `json:"content"`
	ProductId int64  `json:"product_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	gorm.Model
}
