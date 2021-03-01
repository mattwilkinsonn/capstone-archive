package auth

import (
	"context"

	"github.com/Zireael13/capstone-archive/server/internal/db"
	"github.com/gin-contrib/sessions"
)

// Creates a new session from the user's database representation.
func CreateSessionFromUser(ctx context.Context, user *db.User) {
	ginCtx := GinContextFromContext(ctx)

	session := sessions.Default(ginCtx)
	session.Set("userId", user.ID)
	err := session.Save()

	if err != nil {
		panic(err)
	}
}

// Clears session from session/context.
func ClearSession(ctx context.Context) {
	ginCtx := GinContextFromContext(ctx)

	session := sessions.Default(ginCtx)
	session.Clear()
	err := session.Save()

	if err != nil {
		panic(err)
	}
}

// Gets UserID out of the session inside the gin context, which is packed inside the gqlgen context. Returns true/false if id is found.
func GetUserIDFromSession(ctx context.Context) (uint, bool) {
	ginCtx := GinContextFromContext(ctx)

	session := sessions.Default(ginCtx)
	userId := session.Get("userId")

	if userId == nil {
		return 0, false
	}

	return userId.(uint), true

}
