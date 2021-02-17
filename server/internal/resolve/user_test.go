package resolve_test

import (
	"database/sql/driver"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Zireael13/capstone-archive/server/internal/db"
	"github.com/Zireael13/capstone-archive/server/internal/db/dbtest"
	"github.com/Zireael13/capstone-archive/server/internal/graph/model"
	. "github.com/Zireael13/capstone-archive/server/internal/resolve"
	"github.com/Zireael13/capstone-archive/server/internal/serve"
	"github.com/stretchr/testify/assert"
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

		hashed := HashPassword(argon, pass)

		got := VerifyPassword(pass, hashed)

		assert.True(t, got, "Verification should be true")

	})

	t.Run("incorrect password is incorrect", func(t *testing.T) {
		originalPass := "hunter2"

		hashed := HashPassword(argon, originalPass)

		wrongPass := "hunter1"

		got := VerifyPassword(wrongPass, hashed)

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
			name: "Invalid Email",
			input: model.Register{
				Username: "matt12",
				Email:    "mattmatt.com",
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
// TODO: Switch database mocking to https://github.com/Selvatico/go-mocket or just get rid of the orm lol
func TestGetUserFromUsernameOrEmail(t *testing.T) {
	orm, mock := dbtest.CreateMockDBClient(t)

	user := db.User{Username: "matt123", Email: "matt@matt.com", Password: "hunter2"}

	mock.ExpectQuery(
		"INSERT INTO \"users\"",
	).WillReturnRows(
		sqlmock.NewRows()
	).WithArgs(
		AnyTime{}, AnyTime{}, nil, user.Username, user.Email, user.Password,
	)
	orm.Create(&user)

	t.Run("Email", func(t *testing.T) {
		mock.ExpectExec("INSERT \"users\"").WillReturnResult(sqlmock.NewResult(1, 1))
		input := "matt@matt.com"
		returned, err := GetUserFromUsernameOrEmail(input, orm)

		assert.Nil(t, err, "err should be nil")

		assert.Equal(t, user, returned, "returned user shoud match user")
	})
}
