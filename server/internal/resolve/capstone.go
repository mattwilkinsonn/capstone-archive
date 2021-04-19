package resolve

import (
	"context"
	"regexp"
	"time"

	"github.com/Zireael13/capstone-archive/server/internal/db"
	"github.com/Zireael13/capstone-archive/server/internal/graph/model"
	"github.com/gofrs/uuid"
	"github.com/gosimple/slug"
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
		Slug:        capstone.Slug,
	}
}

// Takes Capstone inputs and creates object in Database
func CreateCapstoneInDB(
	ctx context.Context,
	Queries *db.Queries,
	title, description, author, semester string,
) (*db.Capstone, error) {
	id, err := uuid.NewV4()

	if err != nil {
		panic(err)
	}

	input := db.CreateCapstoneParams{
		ID:          id,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Title:       title,
		Description: description,
		Author:      author,
		Semester:    semester,
		Slug:        slug.Make(title),
	}

	capstone, err := Queries.CreateCapstone(ctx, input)

	return &capstone, err
}

// very dumb function right now. Need to add a way to return errors in capstone graphql schema
func HandleCreateCapstoneErr(err error) error {
	return err
}

// Gets list of most recent capstones using cursor pagination.
func GetCapstones(ctx context.Context,
	Queries *db.Queries, limit int, cursor *int) ([]db.Capstone, error) {

	var capstones []db.Capstone
	var err error

	if cursor != nil {
		cursor_time := time.Unix(int64(*cursor), 0)
		params := db.GetCapstonesWithCursorParams{CreatedAt: cursor_time, Limit: int32(limit)}
		capstones, err = Queries.GetCapstonesWithCursor(ctx, params)
	} else {
		capstones, err = Queries.GetCapstones(ctx, int32(limit))
	}

	return capstones, err
}

func GetCapstoneById(ctx context.Context,
	Queries *db.Queries, id string) (*db.Capstone, error) {

	uid, err := uuid.FromString(id)

	if err != nil {
		return nil, err
	}

	capstone, err := Queries.GetCapstoneById(ctx, uid)

	return &capstone, err
}

var whitespace = regexp.MustCompile(`\s+`)

// Searches capstones with Postgres' full text search. Uses LIMIT/OFFSET pagination. Doing LIMIT/OFFSET is probably slow but good enough for use case
func SearchCapstones(
	ctx context.Context,
	Queries *db.Queries,
	query string,
	limit int,
	offset *int,
) ([]db.Capstone, error) {

	var capstones []db.Capstone
	var err error

	query = whitespace.ReplaceAllString(query, "&")

	if offset != nil {
		capstones, err = Queries.SearchCapstones(
			ctx,
			db.SearchCapstonesParams{ToTsquery: query, Limit: int32(limit), Offset: int32(*offset)},
		)
	} else {
		capstones, err = Queries.SearchCapstones(
			ctx,
			db.SearchCapstonesParams{ToTsquery: query, Limit: int32(limit), Offset: 0},
		)
	}

	return capstones, err
}

// Iterates over a slice of DB capstones and creates GQL ones.
func CreateGraphCapstoneSlice(capstones []db.Capstone) (gqlCapstones []*model.Capstone) {
	for _, capstone := range capstones {
		gqlCapstones = append(gqlCapstones, CreateGraphCapstone(&capstone))
	}

	return gqlCapstones
}

func GetCapstoneByTitle(
	ctx context.Context,
	Queries *db.Queries,
	title string,
) (*db.Capstone, error) {
	capstone, err := Queries.GetCapstoneByTitle(ctx, title)
	return &capstone, err
}

func GetCapstoneBySlug(ctx context.Context,
	Queries *db.Queries, slug string) (*db.Capstone, error) {
	capstone, err := Queries.GetCapstoneBySlug(ctx, slug)

	return &capstone, err
}
