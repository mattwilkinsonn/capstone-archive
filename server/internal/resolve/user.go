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

func UIntToString(id uint) string {
	return strconv.FormatUint(uint64(id), 10)
}

func HashPassword(argon *argon2.Config, password string) string {
	hashed, err := argon.HashEncoded([]byte(password))

	if err != nil {
		panic(err)
	}

	return string(hashed)
}

func VerifyPassword(password string, hashed string) bool {

	ok, err := argon2.VerifyEncoded([]byte(password), []byte(hashed))

	if err != nil {
		panic(err)
	}

	return ok
}

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

func CreateUserErr(field, message string) *model.UserError {
	return &model.UserError{Field: field, Message: message}
}

func CreateUserResponseErr(userError *model.UserError) *model.UserResponse {
	return &model.UserResponse{Error: userError}
}

func IsEmail(usernameOrEmail string) bool {
	return strings.Contains(usernameOrEmail, "@")
}

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

func CreateUserResponse(user *db.User) *model.UserResponse {
	return &model.UserResponse{
		User: &model.User{
			ID:        UIntToString(user.ID),
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format(time.UnixDate),
			UpdatedAt: user.UpdatedAt.Format(time.UnixDate),
		}}
}

func CreateUserInDB(DB *gorm.DB, username, email, password string) (*db.User, error) {
	user := db.User{Username: username, Email: email, Password: password}

	res := DB.Create(&user)

	return &user, res.Error
}

func HandleCreateUserErr(err error) *model.UserResponse {
	// TODO: split email and username errors
	if strings.Contains(err.Error(), "23505") {
		return &model.UserResponse{
			Error: &model.UserError{
				Field:   "Email/Username",
				Message: "Email/Username already taken",
			},
		}
	}

	panic(err)
}

func HandleInvalidLogin() *model.UserResponse {
	return &model.UserResponse{
		Error: &model.UserError{Field: "None", Message: "Invalid Login"},
	}
}
