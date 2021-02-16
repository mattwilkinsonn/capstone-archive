package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Zireael13/capstone-archive/server/internal/db"
	"github.com/Zireael13/capstone-archive/server/internal/helpers"

	"github.com/Zireael13/capstone-archive/server/internal/graph/generated"
	"github.com/Zireael13/capstone-archive/server/internal/graph/model"
)

func (r *mutationResolver) CreateCapstone(
	ctx context.Context,
	input model.NewCapstone,
) (*model.Capstone, error) {
	capstone := db.Capstone{
		Title:       input.Title,
		Description: input.Description,
		Author:      input.Author,
	}

	r.DB.Create(&capstone)

	graphCapstone := model.Capstone{
		ID:          helpers.UIntToString(capstone.ID),
		Title:       capstone.Title,
		Description: capstone.Description,
		Author:      capstone.Author,
		CreatedAt:   capstone.CreatedAt.Format(time.UnixDate),
		UpdatedAt:   capstone.UpdatedAt.Format(time.UnixDate),
	}

	return &graphCapstone, nil

}

func (r *mutationResolver) Register(
	ctx context.Context,
	input model.Register,
) (*model.UserResponse, error) {

	ok, userErr := helpers.ValidateRegister(input)
	if !ok {
		return &model.UserResponse{Error: userErr}, nil
	}

	hashed := helpers.HashPassword(r.Argon, input.Password)

	user := db.User{Username: input.Username, Email: input.Email, Password: hashed}

	result := r.DB.Create(&user)

	if result.Error != nil {

		// TODO: split email and username errors
		if strings.Contains(result.Error.Error(), "23505") {
			return &model.UserResponse{
				Error: &model.UserError{
					Field:   "Email/Username",
					Message: "Email/Username already taken",
				},
			}, nil
		}

		panic(result.Error)
	}

	userResponse := model.UserResponse{
		User: &model.User{
			ID:        helpers.UIntToString(user.ID),
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format(time.UnixDate),
			UpdatedAt: user.UpdatedAt.Format(time.UnixDate),
		},
	}

	return &userResponse, nil

}

func (r *mutationResolver) Login(
	ctx context.Context,
	input model.Login,
) (*model.UserResponse, error) {
	panic(fmt.Errorf("not implemented"))
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
