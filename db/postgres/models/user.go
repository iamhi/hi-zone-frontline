package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Uuid     string `gorm:"unique" json:"uuid"`
	Username string `gorm:"unique" json:"username"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Roles    string `gorm:"default:Basic" json:"roles"`
}
