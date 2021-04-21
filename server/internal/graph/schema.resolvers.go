package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Zireael13/capstone-archive/server/internal/auth"
	"github.com/Zireael13/capstone-archive/server/internal/graph/generated"
	"github.com/Zireael13/capstone-archive/server/internal/graph/model"
	. "github.com/Zireael13/capstone-archive/server/internal/resolve"
	"github.com/adam-lavrik/go-imath/ix"
)

func (r *mutationResolver) CreateCapstone(
	ctx context.Context,
	input model.NewCapstone,
) (*model.Capstone, error) {

	session, ok := auth.GetUserSession(ctx)
	if !ok {
		return nil, errors.New("Not logged in")
	}

	ok = CheckAdminRole(session.Role)
	if !ok {
		return nil, errors.New("Not admin role")
	}

	capstone, err := CreateCapstoneInDB(
		ctx,
		r.Queries,
		CreateCapstoneInDBInput(input),
	)

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

	user, err := CreateUserInDB(ctx, r.Queries, r.Argon, CreateUserInDBInput(input))

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
	user, err := GetUserFromUsernameOrEmail(ctx, r.Queries, input.UsernameOrEmail)
	if err != nil {
		return HandleInvalidLogin(), nil
	}

	ok, err := VerifyPassword(input.Password, user.Password)
	// leaving here because might want to handle this error later
	if err != nil {
		log.Fatalf("%v", err)
		// fmt.Errorf()
		// panic(err)
	}

	if !ok {
		return HandleInvalidLogin(), nil
	}

	userResponse := CreateUserResponse(user)

	auth.CreateSessionFromUser(ctx, user)

	return userResponse, nil
}

func (r *mutationResolver) Logout(ctx context.Context) (bool, error) {
	auth.ClearSession(ctx)
	return true, nil
}

func (r *queryResolver) SearchCapstones(
	ctx context.Context,
	query string,
	limit int,
	offset *int,
) (*model.PaginatedCapstones, error) {
	realLimit := ix.Min(limit, 50)

	capstones, err := SearchCapstones(ctx, r.Queries, query, realLimit+1, offset)
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

func (r *queryResolver) Capstones(
	ctx context.Context,
	limit int,
	cursor *int,
) (*model.PaginatedCapstones, error) {
	realLimit := ix.Min(limit, 50)

	capstones, err := GetCapstones(ctx, r.Queries, realLimit+1, cursor)
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

func (r *queryResolver) CapstoneByID(ctx context.Context, id string) (*model.Capstone, error) {
	capstone, err := GetCapstoneById(ctx, r.Queries, id)

	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, nil
	}

	gqlCapstone := CreateGraphCapstone(capstone)

	return gqlCapstone, nil
}

func (r *queryResolver) CapstoneBySlug(ctx context.Context, slug string) (*model.Capstone, error) {
	capstone, err := GetCapstoneBySlug(ctx, r.Queries, slug)

	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, nil
	}

	gqlCapstone := CreateGraphCapstone(capstone)

	return gqlCapstone, nil
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	session, ok := auth.GetUserSession(ctx)
	if !ok {
		return nil, nil
	}

	user, err := GetUserFromID(ctx, r.Queries, session.ID)
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
