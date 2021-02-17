package resolve

import (
	"time"

	"github.com/Zireael13/capstone-archive/server/internal/db"
	"github.com/Zireael13/capstone-archive/server/internal/graph/model"
	"gorm.io/gorm"
)

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

func CreateCapstoneInDB(DB *gorm.DB, title, description, author string) (*db.Capstone, error) {
	capstone := db.Capstone{
		Title:       title,
		Description: description,
		Author:      author,
	}

	res := DB.Create(&capstone)

	return &capstone, res.Error
}

func HandleCreateCapstoneErr(err error) {
	panic(err)
}
