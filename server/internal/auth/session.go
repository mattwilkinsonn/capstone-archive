package auth

import (
	"context"

	"github.com/Zireael13/capstone-archive/server/internal/db"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// Creates a new session from the user's database representation.
func CreateSessionFromUser(ginCtx *gin.Context, user *db.User) {

	session := sessions.Default(ginCtx)
	session.Set("userId", user.ID.String())
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
func GetUserIDFromSession(ctx context.Context) (uuid.UUID, bool) {
	ginCtx := GinContextFromContext(ctx)

	session := sessions.Default(ginCtx)
	userId := session.Get("userId")

	if userId == nil {
		return uuid.Nil, false
	}

	uid, err := uuid.FromString(userId.(string))
	if err != nil {
		panic(err)
	}

	return uid, true

}
