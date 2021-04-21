package auth

import (
	"context"

	"github.com/Zireael13/capstone-archive/server/internal/db"
	"github.com/Zireael13/capstone-archive/server/internal/pack"
	"github.com/gin-contrib/sessions"
	"github.com/gofrs/uuid"
)

type UserSession struct {
	ID   uuid.UUID
	Role db.UserRole
}

// Creates a new session from the user's database representation.
func CreateSessionFromUser(ctx context.Context, user *db.User) {
	ginCtx := pack.GinContextFromContext(ctx)

	session := sessions.Default(ginCtx)
	session.Set("ID", user.ID.String())
	session.Set("Role", string(user.Role))
	err := session.Save()

	if err != nil {
		panic(err)
	}
}

// Clears session from session/context.
func ClearSession(ctx context.Context) {
	ginCtx := pack.GinContextFromContext(ctx)

	session := sessions.Default(ginCtx)
	session.Clear()
	err := session.Save()

	if err != nil {
		panic(err)
	}
}

// Gets UserID out of the session inside the gin context, which is packed inside the gqlgen context. Returns true/false if session is present.
func GetUserSession(ctx context.Context) (UserSession, bool) {
	ginCtx := pack.GinContextFromContext(ctx)

	session := sessions.Default(ginCtx)

	id := session.Get("ID")
	role := session.Get("Role")

	if id == nil {
		return UserSession{}, false
	}

	user_role := db.UserRole("")
	err := user_role.Scan(role)
	if err != nil {
		panic(err)
	}

	uid, err := uuid.FromString(id.(string))
	if err != nil {
		panic(err)
	}

	return UserSession{ID: uid, Role: user_role}, true

}
