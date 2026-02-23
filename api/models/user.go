package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email" gorm:"unique"`
	Password string `json:"password" validate:"required"`
	Phone    string `json:"phone" validate:"required,numeric,min=10,max=10" gorm:"unique"`
	Role     string `json:"role" gorm:"default:'user'"`
	IsActive bool   `json:"is_active" gorm:"default:true"`
	Salt     string `json:"-"`
}
