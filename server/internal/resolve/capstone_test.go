package resolve_test

import (
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
		Base:        db.Base{ID: id, CreatedAt: now, UpdatedAt: now},
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
	orm, mock := dbtest.CreateMockDBClient(t)

	title := "Capstone Archive"
	desc := "Archive for capstone projects"
	author := "Matt Wilkinson"
	semester := "Fall 2019"

	// capstone = &db.Capstone{Title: title, Description: desc, Author: author}

	mock.ExpectQuery(
		regexp.QuoteMeta(`INSERT INTO "capstones"`),
	).WithArgs(AnyTime{}, AnyTime{}, nil, title, desc, author, semester).WillReturnRows(
		mock.NewRows([]string{"id"}).AddRow(1),
	)

	capstone, err := CreateCapstoneInDB(orm, title, desc, author, semester)

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
	orm, mock := dbtest.CreateMockDBClient(t)

	// now := time.Now().Unix()
	limit := 3

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "capstones"`)).
		WithArgs().
		WillReturnRows(mock.NewRows([]string{"id"}).AddRow(1).AddRow(2).AddRow(3))

	capstones, err := GetCapstones(orm, limit, nil)

	require.Nil(t, err)
	assert.Len(t, capstones, limit)
	assert.NotNil(t, capstones)
	assert.Nil(t, mock.ExpectationsWereMet(), "all mock expectations should be met")

}

func TestGetCapstonesWithCursor(t *testing.T) {
	orm, mock := dbtest.CreateMockDBClient(t)

	now := int(time.Now().Unix())
	limit := 3

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "capstones"`)).
		WithArgs(AnyTime{}).
		WillReturnRows(mock.NewRows([]string{"id"}).AddRow(1).AddRow(2).AddRow(3))

	capstones, err := GetCapstones(orm, limit, &now)

	require.Nil(t, err)
	assert.Len(t, capstones, limit)
	assert.NotNil(t, capstones)
	assert.Nil(t, mock.ExpectationsWereMet(), "all mock expectations should be met")

}

func TestCreateGraphCapstoneSlice(t *testing.T) {
	now := time.Now()
	formattedNow := int(now.Unix())

	title := "Capstone Archive"
	desc := "Archive for capstone projects"
	author := "Matt Wilkinson"

	id, _ := uuid.NewV4()

	input := []*db.Capstone{
		{Title: title,
			Description: desc,
			Author:      author,
			Base:        db.Base{ID: id, CreatedAt: now, UpdatedAt: now}},
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
