package resolve

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/Zireael13/capstone-archive/server/internal/db"
	"github.com/Zireael13/capstone-archive/server/internal/graph/model"
	"github.com/gofrs/uuid"
	"github.com/matthewhartstonge/argon2"
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
	case !strings.Contains(input.Email, "@"), len(input.Email) <= 7:
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
func GetUserFromUsernameOrEmail(ctx context.Context,
	Queries *db.Queries, usernameOrEmail string) (*db.User, error) {

	var user db.User
	var err error

	if IsEmail(usernameOrEmail) {
		user, err = Queries.GetUserByEmail(ctx, usernameOrEmail)
	} else {
		user, err = Queries.GetUserByUsername(ctx, usernameOrEmail)
	}

	return &user, err
}

// Transforms DB/ORM User to GraphQL UserResponse
func CreateUserResponse(user *db.User) *model.UserResponse {
	return &model.UserResponse{
		User: DBToGQLUser(user)}
}

func DBToGQLUser(user *db.User) *model.User {
	return &model.User{
		ID:        user.ID.String(),
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: int(user.CreatedAt.Unix()),
		UpdatedAt: int(user.UpdatedAt.Unix()),
		Role:      model.Role(user.Role),
	}
}

// takes User inputs and creates User in DB, returns User object and error if failed
func CreateUserInDB(ctx context.Context,
	Queries *db.Queries, username, email, password string) (*db.User, error) {
	id, err := uuid.NewV4()

	if err != nil {
		panic(err)
	}

	params := db.CreateUserParams{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Username:  username,
		Email:     email,
		Password:  password,
	}

	user, err := Queries.CreateUser(ctx, params)

	return &user, err
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

// Get the user from their ID
func GetUserFromID(ctx context.Context,
	Queries *db.Queries, id uuid.UUID) (*db.User, error) {

	user, err := Queries.GetUserById(ctx, id)

	return &user, err
}

func SetUserToAdmin(ctx context.Context, Queries *db.Queries, id uuid.UUID) (*db.User, error) {
	user, err := Queries.UpdateUserRole(
		ctx,
		db.UpdateUserRoleParams{Role: db.UserRoleADMIN, ID: id},
	)

	return &user, err
}
