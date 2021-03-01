package resolve_test

import (
	"database/sql/driver"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/Zireael13/capstone-archive/server/internal/db"
	"github.com/Zireael13/capstone-archive/server/internal/db/dbtest"
	"github.com/Zireael13/capstone-archive/server/internal/graph/model"
	. "github.com/Zireael13/capstone-archive/server/internal/resolve"
	"github.com/Zireael13/capstone-archive/server/internal/serve"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestUIntToString(t *testing.T) {
	var id uint = 32
	want := "32"

	got := UIntToString(id)

	assert.Equal(t, got, want)

}

func TestHashAndVerifyPassword(t *testing.T) {
	argon := serve.CreateArgon()
	t.Run("correct password hash and dehash", func(t *testing.T) {
		pass := "hunter2"

		hashed, err := HashPassword(argon, pass)

		require.Nil(t, err)

		got, err := VerifyPassword(pass, hashed)

		require.Nil(t, err)

		assert.True(t, got, "Verification should be true")

	})

	t.Run("incorrect password is incorrect", func(t *testing.T) {
		originalPass := "hunter2"

		hashed, err := HashPassword(argon, originalPass)

		require.Nil(t, err)

		wrongPass := "hunter1"

		got, err := VerifyPassword(wrongPass, hashed)

		require.Nil(t, err)

		assert.False(t, got, "Verification should be false")

	})

}

func TestValidateRegister(t *testing.T) {
	t.Run("Correct Registration", func(t *testing.T) {
		input := model.Register{Username: "matt12", Email: "matt@matt.com", Password: "hunter2"}

		ok, res := ValidateRegister(input)

		assert.True(t, ok, "Register should be true")

		assert.Equal(
			t,
			&model.UserResponse{Error: &model.UserError{Field: "none", Message: "none"}},
			res,
			"Response should be no error",
		)
	})

	validateRegisterTests := []struct {
		name     string
		input    model.Register
		valid    bool
		response *model.UserResponse
	}{
		{
			name: "Correct",
			input: model.Register{
				Username: "matt12",
				Email:    "matt@matt.com",
				Password: "hunter2",
			},
			valid:    true,
			response: CreateUserResponseErr(CreateUserErr("none", "none")),
		},
		{
			name: "Invalid Email - No @",
			input: model.Register{
				Username: "matt12",
				Email:    "mattmatt.com",
				Password: "hunter2",
			},
			valid:    false,
			response: CreateUserResponseErr(CreateUserErr("Email", "Invalid Email")),
		},
		{
			name: "Invalid Email - Too Short",
			input: model.Register{
				Username: "matt12",
				Email:    "mat@",
				Password: "hunter2",
			},
			valid:    false,
			response: CreateUserResponseErr(CreateUserErr("Email", "Invalid Email")),
		},
		{
			name: "Username too short",
			input: model.Register{
				Username: "mat",
				Email:    "matt@matt.com",
				Password: "hunter2",
			},
			valid: false,
			response: CreateUserResponseErr(
				CreateUserErr("Username", "Username must be longer than 4 characters"),
			),
		},
		{
			name: "Password too short",
			input: model.Register{
				Username: "matt12",
				Email:    "matt@matt.com",
				Password: "hunter",
			},
			valid: false,
			response: CreateUserResponseErr(
				CreateUserErr("Password", "Password must be longer than 6 characters"),
			),
		},
	}
	for _, tt := range validateRegisterTests {
		t.Run(tt.name, func(t *testing.T) {
			ok, res := ValidateRegister(tt.input)

			assert.Equal(t, tt.valid, ok)

			assert.Equal(t, tt.response, res)
		})
	}

}

func TestCreateUserErr(t *testing.T) {
	field := "Email"
	message := "Invalid Email"
	want := &model.UserError{Field: field, Message: message}

	got := CreateUserErr(field, message)

	assert.Equal(t, want, got, "Expected UserError to be created correctly")
}

func TestCreateUserResponseErr(t *testing.T) {
	err := CreateUserErr("Username", "Invalid Username")

	want := &model.UserResponse{Error: err}

	got := CreateUserResponseErr(err)

	assert.Equal(t, want, got, "Expected UserResponse to be created correctly")
}

func TestIsEmail(t *testing.T) {
	t.Run("Email", func(t *testing.T) {
		email := "matt@matt.com"
		got := IsEmail(email)
		assert.True(t, got, "Expected true for email")
	})

	t.Run("Username", func(t *testing.T) {
		username := "matt12346"
		got := IsEmail(username)
		assert.False(t, got, "Expected false for username")
	})
}

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestGetUserFromUsernameOrEmail(t *testing.T) {
	orm, mock := dbtest.CreateMockDBClient(t)

	user := &db.User{Username: "matt123", Email: "matt@matt.com", Password: "hunter2"}

	mock.ExpectQuery(
		"INSERT INTO \"users\"",
	).WithArgs(
		AnyTime{}, AnyTime{}, nil, user.Username, user.Email, user.Password,
	).WillReturnRows(mock.NewRows([]string{"id"}))
	orm.Create(&user)

	t.Run("Email", func(t *testing.T) {
		mock.ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "users" WHERE email = $1`),
		).WithArgs(
			user.Email,
		).WillReturnRows(mock.NewRows([]string{"id",
			"username",
			"email",
			"created_at",
			"updated_at",
			"deleted_at",
			"password"}).AddRow(user.ID,
			user.Username,
			user.Email,
			user.CreatedAt,
			user.UpdatedAt,
			user.DeletedAt,
			user.Password))

		input := "matt@matt.com"
		returned, err := GetUserFromUsernameOrEmail(input, orm)

		assert.Nil(t, err, "err should be nil")

		assert.Equal(t, user, returned, "returned user shoud match user")
	})

	t.Run("Username", func(t *testing.T) {
		mock.ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "users" WHERE username = $1`),
		).WithArgs(
			user.Username,
		).WillReturnRows(mock.NewRows([]string{"id",
			"username",
			"email",
			"created_at",
			"updated_at",
			"deleted_at",
			"password"}).AddRow(user.ID,
			user.Username,
			user.Email,
			user.CreatedAt,
			user.UpdatedAt,
			user.DeletedAt,
			user.Password))

		input := user.Username
		returned, err := GetUserFromUsernameOrEmail(input, orm)

		assert.Nil(t, err, "err should be nil")

		assert.Equal(t, user, returned, "returned user shoud match user")

	})

	assert.Nil(t, mock.ExpectationsWereMet(), "all mock expectations should be met")
}

func TestCreateUserResponse(t *testing.T) {

	now := time.Now()
	formattedNow := now.Format(time.UnixDate)
	user := &db.User{
		Model:    gorm.Model{ID: 1, CreatedAt: now, UpdatedAt: now},
		Username: "zireael",
		Email:    "zir@gmail.com",
		Password: "hunter2",
	}

	want := &model.UserResponse{
		User: &model.User{
			ID:        1,
			Username:  "zireael",
			Email:     "zir@gmail.com",
			CreatedAt: formattedNow,
			UpdatedAt: formattedNow,
		},
	}

	got := CreateUserResponse(user)

	assert.Equal(t, want, got, "Returned user response should mirror wanted one")
}

func TestCreateUserInDB(t *testing.T) {
	orm, mock := dbtest.CreateMockDBClient(t)
	username := "Zireael"
	email := "zir@gmail.com"
	password := "hunter2"

	mock.ExpectQuery(
		regexp.QuoteMeta(`INSERT INTO "users"`),
	).WithArgs(
		AnyTime{},
		AnyTime{},
		nil,
		username,
		email,
		password,
	).WillReturnRows(mock.NewRows([]string{"id"}).AddRow(1))

	got, err := CreateUserInDB(orm, username, email, password)

	assert.Nil(t, err, "should be no err")

	assert.Equal(t, email, got.Email, "returned user email should be wanted user email")

	assert.Nil(t, mock.ExpectationsWereMet(), "all mock expectations should be met")
}

func TestHandleCreateUserErr(t *testing.T) {

	t.Run("Duplicate Username", func(t *testing.T) {
		err := errors.New(
			`ERROR: duplicate key value violates unique constraint "users_username_key" (SQLSTATE 23505)`,
		)

		want := CreateUserResponseErr(CreateUserErr("Username", "Username already taken"))

		got, returnedErr := HandleCreateUserErr(err)

		assert.Equal(t, want, got, "Duplicate Username Response")
		assert.Nil(t, returnedErr)
	})

	t.Run("Duplicate Email", func(t *testing.T) {
		err := errors.New(
			`ERROR: duplicate key value violates unique constraint "users_email_key" (SQLSTATE 23505)`,
		)

		want := CreateUserResponseErr(CreateUserErr("Email", "Email already taken"))

		got, returnedErr := HandleCreateUserErr(err)

		assert.Equal(t, want, got, "Duplicate Email Response")
		assert.Nil(t, returnedErr)
	})

	t.Run("Unhandled Error", func(t *testing.T) {
		err := errors.New(`ERROR: datetime_field_overflow "users_created_at" (SQLSTATE 22008)`)

		want := &model.UserResponse{}

		got, returnedErr := HandleCreateUserErr(err)

		assert.Equal(t, want, got, "Response should be blank")
		assert.Error(t, returnedErr, "Unhandled Err should be returned")

	})
}

func TestHandleInvalidLogin(t *testing.T) {
	want := &model.UserResponse{Error: &model.UserError{Field: "None", Message: "Invalid Login"}}

	got := HandleInvalidLogin()

	assert.Equal(t, want, got)
}
