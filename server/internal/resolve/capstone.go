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
		ID:          int(capstone.ID),
		Title:       capstone.Title,
		Description: capstone.Description,
		Author:      capstone.Author,
		CreatedAt:   int(capstone.CreatedAt.Unix()),
		UpdatedAt:   int(capstone.UpdatedAt.Unix()),
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

func GetCapstones(DB *gorm.DB, number int, cursor *int) (capstones []*db.Capstone, err error) {
	var res *gorm.DB

	if cursor != nil {
		res = DB.Where(
			"created_at < ?",
			time.Unix(int64(*cursor), 0),
		).Order(
			"created_at DESC",
		).Limit(
			number,
		).Find(
			&capstones,
		)
	} else {
		res = DB.Limit(number).Order("created_at DESC").Find(&capstones)
	}

	return capstones, res.Error
}

func CreateGraphCapstoneSlice(capstones []*db.Capstone) (gqlCapstones []*model.Capstone) {
	for _, capstone := range capstones {
		gqlCapstones = append(gqlCapstones, CreateGraphCapstone(capstone))
	}

	return gqlCapstones
}
