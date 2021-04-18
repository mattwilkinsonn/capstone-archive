package db

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// Base contains common columns for all tables.
type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID, _ = uuid.NewV4()

	return
}

type User struct {
	Base
	Username string `gorm:"unique;not null;"`
	Email    string `gorm:"unique;not null;"`
	Password string `gorm:"not null;"`
	Role     string `gorm:"not null;default:USER;"`
}

type Capstone struct {
	Base
	Title       string `gorm:"not null;"faker:"Sentence"`
	Description string `gorm:"not null;"faker:"Sentence"`
	Author      string `gorm:"not null;"faker:"FirstName"`
	Semester    string `gorm:"not null;"faker:"Word"`
}
