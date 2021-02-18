package resolve

import (
	"strconv"
	"strings"
	"time"

	"github.com/Zireael13/capstone-archive/server/internal/db"
	"github.com/Zireael13/capstone-archive/server/internal/graph/model"
	"github.com/matthewhartstonge/argon2"
	"gorm.io/gorm"
)

// Converts UInt type to String
func UIntToString(id uint) string {
	return strconv.FormatUint(uint64(id), 10)
}

// Hash password
func HashPassword(argon *argon2.Config, password string) (string, error) {
	hashed, err := argon.HashEncoded([]byte(password))

	return string(hashed), err
}

// Dehash password and check it against hashed
func VerifyPassword(password string, hashed string) (bool, error) {

	ok, err := argon2.VerifyEncoded([]byte(password), []byte(hashed))

	return ok, err
}

// Input validation on Register
func ValidateRegister(input model.Register) (bool, *model.UserResponse) {
	var err *model.UserError
	var ok bool

	switch {
	case !strings.Contains(input.Email, "@"):
		ok = false
		err = CreateUserErr("Email", "Invalid Email")
	case len(input.Username) <= 4:
		ok = false
		err = CreateUserErr("Username", "Username must be longer than 4 characters")
	case len(input.Password) <= 6:
		ok = false
		err = CreateUserErr("Password", "Password must be longer than 6 characters")
	default:
		ok = true
		err = CreateUserErr("none", "none")
	}

	res := CreateUserResponseErr(err)

	return ok, res
}

// Wrapper for creating GQL UserErrors
func CreateUserErr(field, message string) *model.UserError {
	return &model.UserError{Field: field, Message: message}
}

// Wrapper for creating GQL UserResponses with an error
func CreateUserResponseErr(userError *model.UserError) *model.UserResponse {
	return &model.UserResponse{Error: userError}
}

// Checks if string is an email.
func IsEmail(usernameOrEmail string) bool {
	return strings.Contains(usernameOrEmail, "@")
}

// Queries DB for User on an ambigious username or email param.
func GetUserFromUsernameOrEmail(usernameOrEmail string, DB *gorm.DB) (db.User, error) {
	var res *gorm.DB
	user := db.User{}

	if IsEmail(usernameOrEmail) {
		res = DB.Where("email = ?", usernameOrEmail).First(&user)
	} else {
		res = DB.Where("username = ?", usernameOrEmail).First(&user)
	}

	return user, res.Error
}

// Transforms DB/ORM User to GraphQL UserResponse
func CreateUserResponse(user *db.User) *model.UserResponse {
	return &model.UserResponse{
		User: &model.User{
			ID:        int(user.ID),
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format(time.UnixDate),
			UpdatedAt: user.UpdatedAt.Format(time.UnixDate),
		}}
}

// takes User inputs and creates User in DB, returns User object and error if failed
func CreateUserInDB(DB *gorm.DB, username, email, password string) (*db.User, error) {
	user := db.User{Username: username, Email: email, Password: password}

	res := DB.Create(&user)

	return &user, res.Error
}

const DuplicateUsernameErr = `ERROR: duplicate key value violates unique constraint "users_username_key" (SQLSTATE 23505)`

const DuplicateEmailErr = `ERROR: duplicate key value violates unique constraint "users_email_key" (SQLSTATE 23505)`

// Handles errors from creating User in DB
func HandleCreateUserErr(err error) (*model.UserResponse, error) {
	switch err.Error() {
	case DuplicateUsernameErr:
		return CreateUserResponseErr(CreateUserErr("Username", "Username already taken")), nil
	case DuplicateEmailErr:
		return CreateUserResponseErr(CreateUserErr("Email", "Email already taken")), nil
	default:
		return &model.UserResponse{}, err
	}
}

// Returns invalid login response
func HandleInvalidLogin() *model.UserResponse {
	return CreateUserResponseErr(CreateUserErr("None", "Invalid Login"))
}
