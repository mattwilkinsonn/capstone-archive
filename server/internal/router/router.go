package router

import (
	"github.com/gin-gonic/gin"
)

func AttachRoutes(g *gin.Engine, gqlHandler gin.HandlerFunc, playHandler gin.HandlerFunc) {
	g.POST("/graphql", gqlHandler)
	g.GET("/", playHandler)
	g.GET("/healthcheck", func(c *gin.Context) {
		c.Status(200)
	})
}
