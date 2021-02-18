package resolve

import (
	"time"

	"github.com/Zireael13/capstone-archive/server/internal/db"
	"github.com/Zireael13/capstone-archive/server/internal/graph/model"
	"gorm.io/gorm"
)

// Transforms DB/ORM Capstone schema into GraphQL schema
func CreateGraphCapstone(capstone *db.Capstone) *model.Capstone {
	return &model.Capstone{
		ID:          UIntToString(capstone.ID),
		Title:       capstone.Title,
		Description: capstone.Description,
		Author:      capstone.Author,
		CreatedAt:   capstone.CreatedAt.Format(time.UnixDate),
		UpdatedAt:   capstone.UpdatedAt.Format(time.UnixDate),
	}
}

// Takes Capstone inputs and creates object in Database
func CreateCapstoneInDB(DB *gorm.DB, title, description, author string) (*db.Capstone, error) {
	capstone := db.Capstone{
		Title:       title,
		Description: description,
		Author:      author,
	}

	res := DB.Create(&capstone)

	return &capstone, res.Error
}

// very dumb function right now. Need to add a way to return errors in capstone graphql schema
func HandleCreateCapstoneErr(err error) error {
	return err
}
