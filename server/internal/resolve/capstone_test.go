package resolve_test

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/Zireael13/capstone-archive/server/internal/db"
	"github.com/Zireael13/capstone-archive/server/internal/db/dbtest"
	"github.com/Zireael13/capstone-archive/server/internal/graph/model"
	. "github.com/Zireael13/capstone-archive/server/internal/resolve"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateGraphCapstone(t *testing.T) {
	now := time.Now()
	formattedNow := int(now.Unix())

	title := "Capstone Archive"
	desc := "Archive for capstone projects"
	author := "Matt Wilkinson"

	id, _ := uuid.NewV4()

	input := &db.Capstone{
		Title:       title,
		Description: desc,
		Author:      author,
		ID:          id,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	want := &model.Capstone{
		ID:          id.String(),
		Title:       title,
		Description: desc,
		Author:      author,
		CreatedAt:   formattedNow,
		UpdatedAt:   formattedNow,
	}

	got := CreateGraphCapstone(input)

	assert.Equal(t, want, got)
}

func TestCreateCapstoneInDB(t *testing.T) {
	queries, mock := dbtest.CreateMockDBClient(t)

	title := "Capstone Archive"
	desc := "Archive for capstone projects"
	author := "Matt Wilkinson"
	semester := "Fall 2019"

	// capstone = &db.Capstone{Title: title, Description: desc, Author: author}

	id, _ := uuid.NewV4()
	mock.ExpectQuery(
		regexp.QuoteMeta(`INSERT INTO capstones`),
	).WithArgs(Any{}, AnyTime{}, AnyTime{}, title, desc, author, semester).WillReturnRows(
		mock.NewRows([]string{"id",
			"created_at",
			"updated_at",
			"deleted_at",
			"title",
			"description",
			"author",
			"semester"}).AddRow(id,
			time.Now(),
			time.Now(),
			nil,
			title,
			desc,
			author,
			semester),
	)

	capstone, err := CreateCapstoneInDB(
		context.Background(),
		queries,
		title,
		desc,
		author,
		semester,
	)

	assert.Equal(t, author, capstone.Author, "Authors should be equal")
	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet(), "all mock expectations should be met")
}

func TestHandleCreateCapstoneErr(t *testing.T) {

	err := errors.New("some err")

	got := HandleCreateCapstoneErr(err)

	assert.Equal(t, err, got)

}

func TestGetCapstones(t *testing.T) {
	queries, mock := dbtest.CreateMockDBClient(t)

	limit := 3
	id1, _ := uuid.NewV4()
	id2, _ := uuid.NewV4()
	id3, _ := uuid.NewV4()

	mock.ExpectQuery(
		regexp.QuoteMeta(
			`SELECT id, created_at, updated_at, deleted_at, title, description, author, semester FROM capstones`,
		),
	).
		WithArgs().
		WillReturnRows(
			mock.NewRows(
				[]string{
					"id",
					"created_at",
					"updated_at",
					"deleted_at",
					"title",
					"description",
					"author",
					"semester",
				},
			).AddRow(
				id1,
				time.Now(),
				time.Now(),
				nil,
				"Lalalala",
				"this is a description",
				"Matt",
				"Fall 2019",
			).AddRow(
				id2,
				time.Now(),
				time.Now(),
				nil,
				"Lalalala",
				"this is a description",
				"Matt",
				"Fall 2019",
			).AddRow(
				id3,
				time.Now(),
				time.Now(),
				nil,
				"Lalalala",
				"this is a description",
				"Matt",
				"Fall 2019",
			),
		)

	capstones, err := GetCapstones(context.Background(), queries, limit, nil)

	require.Nil(t, err)
	assert.Len(t, capstones, limit)
	assert.NotNil(t, capstones)
	assert.Equal(t, capstones[0].ID, id1)
	assert.Nil(t, mock.ExpectationsWereMet(), "all mock expectations should be met")

}

func TestGetCapstonesWithCursor(t *testing.T) {
	queries, mock := dbtest.CreateMockDBClient(t)

	now := int(time.Now().Unix())
	limit := 3

	id1, _ := uuid.NewV4()
	id2, _ := uuid.NewV4()
	id3, _ := uuid.NewV4()

	mock.ExpectQuery(
		regexp.QuoteMeta(
			`SELECT id, created_at, updated_at, deleted_at, title, description, author, semester FROM capstones`,
		),
	).
		WithArgs(
			AnyTime{},
			limit,
		).
		WillReturnRows(
			mock.NewRows(
				[]string{
					"id",
					"created_at",
					"updated_at",
					"deleted_at",
					"title",
					"description",
					"author",
					"semester",
				},
			).AddRow(
				id1,
				time.Now(),
				time.Now(),
				nil,
				"Lalalala",
				"this is a description",
				"Matt",
				"Fall 2019",
			).AddRow(
				id2,
				time.Now(),
				time.Now(),
				nil,
				"Lalalala",
				"this is a description",
				"Matt",
				"Fall 2019",
			).AddRow(
				id3,
				time.Now(),
				time.Now(),
				nil,
				"Lalalala",
				"this is a description",
				"Matt",
				"Fall 2019",
			),
		)

	capstones, err := GetCapstones(context.Background(), queries, limit, &now)

	require.Nil(t, err)
	assert.Len(t, capstones, limit)
	assert.NotNil(t, capstones)
	assert.Equal(t, capstones[0].ID, id1)
	assert.Nil(t, mock.ExpectationsWereMet(), "all mock expectations should be met")

}

func TestCreateGraphCapstoneSlice(t *testing.T) {
	now := time.Now()
	formattedNow := int(now.Unix())

	title := "Capstone Archive"
	desc := "Archive for capstone projects"
	author := "Matt Wilkinson"

	id, _ := uuid.NewV4()

	input := []db.Capstone{
		{Title: title,
			Description: desc,
			Author:      author,
			ID:          id, CreatedAt: now, UpdatedAt: now},
	}

	want := []*model.Capstone{
		{
			ID:          id.String(),
			Title:       title,
			Description: desc,
			Author:      author,
			CreatedAt:   formattedNow,
			UpdatedAt:   formattedNow,
		},
	}

	got := CreateGraphCapstoneSlice(input)

	assert.Equal(t, want, got)

}
