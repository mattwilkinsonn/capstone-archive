package helpers

import (
	"strconv"
	"strings"

	"github.com/Zireael13/capstone-archive/server/internal/graph/model"
	"github.com/matthewhartstonge/argon2"
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

func ValidateRegister(input model.Register) (bool, *model.UserError) {
	if !strings.Contains(input.Email, "@") {
		return false, &model.UserError{Field: "Email", Message: "Invalid Email"}
	}

	if len(input.Username) < 4 {
		return false, &model.UserError{
			Field:   "Username",
			Message: "Username must be longer than 4 characters",
		}
	}

	if len(input.Password) < 6 {
		return false, &model.UserError{
			Field:   "Password",
			Message: "Password must be longer than 6 characters",
		}
	}

	return true, &model.UserError{Field: "none", Message: "none"}
}
