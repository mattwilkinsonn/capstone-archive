package auth

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

type GinContextKey string

// Packs the Gin Context into the regular context.
func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		k := GinContextKey("gin")
		ctx := context.WithValue(c.Request.Context(), k, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// Extracts the Gin context from the packed Gqlgen context.
func GinContextFromContext(ctx context.Context) *gin.Context {
	ginContext := ctx.Value(GinContextKey("gin"))
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		panic(err)
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		panic(err)
	}
	return gc
}
