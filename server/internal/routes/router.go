package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
	"gorm.io/gorm"
)

func AttachRoutes(g *gin.Engine, orm *gorm.DB, argon *argon2.Config) {
	g.POST("/graphql", GraphQLHandler(orm, argon))
	g.GET("/", PlaygroundHandler())
}
