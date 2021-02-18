package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Zireael13/capstone-archive/server/internal/graph/generated"
	"github.com/Zireael13/capstone-archive/server/internal/graph/model"
	. "github.com/Zireael13/capstone-archive/server/internal/resolve"
)

func (r *mutationResolver) CreateCapstone(
	ctx context.Context,
	input model.NewCapstone,
) (*model.Capstone, error) {
	// TODO: input validation on capstone adds?

	capstone, err := CreateCapstoneInDB(r.DB, input.Title, input.Description, input.Author)
	if err != nil {
		err = HandleCreateCapstoneErr(err)
		panic(err)
	}

	graphCapstone := CreateGraphCapstone(capstone)

	return graphCapstone, nil
}

func (r *mutationResolver) Register(
	ctx context.Context,
	input model.Register,
) (*model.UserResponse, error) {
	ok, res := ValidateRegister(input)
	if !ok {
		return res, nil
	}

	hashed, err := HashPassword(r.Argon, input.Password)
	if err != nil {
		panic(err)
	}

	user, err := CreateUserInDB(r.DB, input.Username, input.Email, hashed)

	if err != nil {
		res, unhandledErr := HandleCreateUserErr(err)
		if unhandledErr != nil {
			panic(unhandledErr)
		}
		return res, nil
	}

	return CreateUserResponse(user), nil
}

func (r *mutationResolver) Login(
	ctx context.Context,
	input model.Login,
) (*model.UserResponse, error) {
	user, err := GetUserFromUsernameOrEmail(input.UsernameOrEmail, r.DB)
	if err != nil {
		return HandleInvalidLogin(), nil
	}

	ok, err := VerifyPassword(input.Password, user.Password)
	if err != nil {
		panic(err)
	}

	if !ok {
		return HandleInvalidLogin(), nil
	}

	userResponse := CreateUserResponse(&user)

	// TODO: implement jwt tokens
	return userResponse, nil
}

func (r *queryResolver) Capstones(ctx context.Context) ([]*model.Capstone, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
