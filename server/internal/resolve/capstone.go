package resolve

import (
	"regexp"
	"time"

	"github.com/Zireael13/capstone-archive/server/internal/db"
	"github.com/Zireael13/capstone-archive/server/internal/graph/model"
	"gorm.io/gorm"
)

// Transforms DB/ORM Capstone schema into GraphQL schema
func CreateGraphCapstone(capstone *db.Capstone) *model.Capstone {
	return &model.Capstone{
		ID:          capstone.ID.String(),
		Title:       capstone.Title,
		Description: capstone.Description,
		Author:      capstone.Author,
		CreatedAt:   int(capstone.CreatedAt.Unix()),
		UpdatedAt:   int(capstone.UpdatedAt.Unix()),
		Semester:    capstone.Semester,
	}
}

// Takes Capstone inputs and creates object in Database
func CreateCapstoneInDB(
	DB *gorm.DB,
	title, description, author, semester string,
) (*db.Capstone, error) {
	capstone := db.Capstone{
		Title:       title,
		Description: description,
		Author:      author,
		Semester:    semester,
	}

	res := DB.Create(&capstone)

	return &capstone, res.Error
}

// very dumb function right now. Need to add a way to return errors in capstone graphql schema
func HandleCreateCapstoneErr(err error) error {
	return err
}

// Gets list of most recent capstones using cursor pagination.
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

func GetCapstoneById(DB *gorm.DB, id uint) (*db.Capstone, error) {

	capstone := db.Capstone{}

	res := DB.Where("id = ?", id).First(&capstone)

	return &capstone, res.Error
}

// Searches capstones with Postgres' full text search. Uses LIMIT/OFFSET pagination. Doing LIMIT/OFFSET is probably slow but good enough for use case
func SearchCapstones(
	DB *gorm.DB,
	query string,
	number int,
	offset *int,
) (capstones []*db.Capstone, err error) {
	var res *gorm.DB

	whitespace := regexp.MustCompile(`\s+`)

	query = whitespace.ReplaceAllString(query, "&")

	if offset != nil {
		sql := "SELECT * FROM (SELECT to_tsvector(c.Title) || to_tsvector(c.Description) || to_tsvector(c.Author) || to_tsvector(c.Semester) as document, * FROM capstones c) capstone WHERE capstone.document @@ to_tsquery('english', ?) LIMIT ? OFFSET ?;"
		res = DB.Raw(sql, query, number, offset).Scan(&capstones)
	} else {
		sql := "SELECT * FROM (SELECT to_tsvector(c.Title) || to_tsvector(c.Description) || to_tsvector(c.Author) || to_tsvector(c.Semester) as document, * FROM capstones c) capstone WHERE capstone.document @@ to_tsquery('english', ?) LIMIT ?"
		res = DB.Raw(sql, query, number).Scan(&capstones)
	}

	return capstones, res.Error
}

// Iterates over a slice of DB capstones and creates GQL ones.
func CreateGraphCapstoneSlice(capstones []*db.Capstone) (gqlCapstones []*model.Capstone) {
	for _, capstone := range capstones {
		gqlCapstones = append(gqlCapstones, CreateGraphCapstone(capstone))
	}

	return gqlCapstones
}
