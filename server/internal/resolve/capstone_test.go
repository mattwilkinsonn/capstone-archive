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
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateGraphCapstone(t *testing.T) {
	now := time.Now()
	formattedNow := now.Format(time.UnixDate)

	title := "Capstone Archive"
	desc := "Archive for capstone projects"
	author := "Matt Wilkinson"

	input := &db.Capstone{
		Title:       title,
		Description: desc,
		Author:      author,
		Model:       gorm.Model{ID: 24, CreatedAt: now, UpdatedAt: now},
	}

	want := &model.Capstone{
		ID:          "24",
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

	// capstone = &db.Capstone{Title: title, Description: desc, Author: author}

	mock.ExpectQuery(
		regexp.QuoteMeta(`INSERT INTO "capstones"`),
	).WithArgs(AnyTime{}, AnyTime{}, nil, title, desc, author).WillReturnRows(
		mock.NewRows([]string{"id"}).AddRow(1),
	)

	capstone, err := CreateCapstoneInDB(orm, title, desc, author)

	assert.Equal(t, author, capstone.Author, "Authors should be equal")
	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet(), "all mock expectations should be met")
}

func TestHandleCreateCapstoneErr(t *testing.T) {

	err := errors.New("some err")

	got := HandleCreateCapstoneErr(err)

	assert.Equal(t, err, got)

}
