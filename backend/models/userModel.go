package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model

	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Password string `json:"password"`
}