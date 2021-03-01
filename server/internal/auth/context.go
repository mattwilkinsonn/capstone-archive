package auth

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

type GinContextKey string

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		k := GinContextKey("gin")
		ctx := context.WithValue(c.Request.Context(), k, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(GinContextKey("gin"))
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}
