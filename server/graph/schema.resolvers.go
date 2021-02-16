package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Zireael13/capstone-archive/server/internal/graph/generated"
	"github.com/Zireael13/capstone-archive/server/internal/graph/model"
)

func (r *mutationResolver) CreateCapstone(ctx context.Context, input model.NewCapstone) (*model.Capstone, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Register(ctx context.Context, input model.Register) (*model.UserResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (*model.UserResponse, error) {
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
