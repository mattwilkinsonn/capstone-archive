package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Zireael13/capstone-archive/server/internal/auth"
	"github.com/Zireael13/capstone-archive/server/internal/graph/generated"
	"github.com/Zireael13/capstone-archive/server/internal/graph/model"
	. "github.com/Zireael13/capstone-archive/server/internal/resolve"
	"github.com/adam-lavrik/go-imath/ix"
)

func (r *mutationResolver) CreateCapstone(ctx context.Context, input model.NewCapstone) (*model.Capstone, error) {
	capstone, err := CreateCapstoneInDB(
		r.DB,
		input.Title,
		input.Description,
		input.Author,
		input.Semester,
	)
	if err != nil {
		err = HandleCreateCapstoneErr(err)
		panic(err)
	}

	graphCapstone := CreateGraphCapstone(capstone)

	return graphCapstone, nil
}

func (r *mutationResolver) Register(ctx context.Context, input model.Register) (*model.UserResponse, error) {
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

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (*model.UserResponse, error) {
	user, err := GetUserFromUsernameOrEmail(input.UsernameOrEmail, r.DB)
	if err != nil {
		return HandleInvalidLogin(), nil
	}

	ok, err := VerifyPassword(input.Password, user.Password)
	// leaving here because might want to handle this error later
	if err != nil {
		panic(err)
	}

	if !ok {
		return HandleInvalidLogin(), nil
	}

	userResponse := CreateUserResponse(user)

	ginCtx := auth.GinContextFromContext(ctx)
	auth.CreateSessionFromUser(ginCtx, user)

	return userResponse, nil
}

func (r *mutationResolver) Logout(ctx context.Context) (bool, error) {
	auth.ClearSession(ctx)
	return true, nil
}

func (r *queryResolver) SearchCapstones(ctx context.Context, query string, limit int, offset *int) (*model.PaginatedCapstones, error) {
	realLimit := ix.Min(limit, 50)

	capstones, err := SearchCapstones(r.DB, query, realLimit+1, offset)
	if err != nil {
		panic(err)
	}

	hasMore := false
	if len(capstones) == realLimit+1 {
		hasMore = true
	} else if len(capstones) < realLimit {
		realLimit = len(capstones)
	}

	gqlCapstones := CreateGraphCapstoneSlice(capstones[0:realLimit])

	paginated := &model.PaginatedCapstones{
		Capstones: gqlCapstones,
		HasMore:   hasMore,
	}

	return paginated, nil
}

func (r *queryResolver) Capstones(ctx context.Context, limit int, cursor *int) (*model.PaginatedCapstones, error) {
	realLimit := ix.Min(limit, 50)

	capstones, err := GetCapstones(r.DB, realLimit+1, cursor)
	if err != nil {
		panic(err)
	}

	hasMore := false
	if len(capstones) == realLimit+1 {
		hasMore = true
	} else if len(capstones) < realLimit {
		realLimit = len(capstones)
	}

	gqlCapstones := CreateGraphCapstoneSlice(capstones[0:realLimit])

	paginated := &model.PaginatedCapstones{
		Capstones: gqlCapstones,
		HasMore:   hasMore,
	}

	return paginated, nil
}

func (r *queryResolver) Capstone(ctx context.Context, id int) (*model.Capstone, error) {
	capstone, err := GetCapstoneById(r.DB, uint(id))

	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, nil
	}

	gqlCapstone := CreateGraphCapstone(capstone)

	return gqlCapstone, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.PublicUser, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	id, ok := auth.GetUserIDFromSession(ctx)
	if !ok {
		return nil, nil
	}

	user, err := GetUserFromID(r.DB, id)
	if err != nil {
		return nil, nil
	}

	gqlUser := DBToGQLUser(user)

	return gqlUser, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
