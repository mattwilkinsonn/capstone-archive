package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null;"`
	Email    string `gorm:"unique;not null;"`
	Password string `gorm:"not null;"`
}

type Capstone struct {
	gorm.Model
	Title       string `gorm:"not null;"`
	Description string `gorm:"not null;"`
	Author      string `gorm:"not null;"`
}
