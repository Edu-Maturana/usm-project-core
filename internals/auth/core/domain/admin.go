package domain

import "gorm.io/gorm"

type Admin struct {
	Username string `gorm:"type:varchar(100);unique_index" json:"username" validate:"required"`
	Email    string `gorm:"type:varchar(100);unique_index" json:"email" validate:"required"`
	Password string `gorm:"type:varchar(100)" json:"password"`
	Active   bool   `gorm:"default:false" json:"active"`
	gorm.Model
}

type User struct {
	Admin
}
