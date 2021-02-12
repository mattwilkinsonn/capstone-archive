package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

type Capstone struct {
	gorm.Model
	Title       string
	Description string
	Author      string
}
