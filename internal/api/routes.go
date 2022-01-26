package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lipeavelar/soccer-bet-api/internal/middlewares"
)

// SetupRoutes sets up the routes for the API
func SetupRoutes(engine *gin.Engine) {
	setupUnauthenticatedRoutes(engine.Group("v1"))
	setupAuthRoutes(engine.Group("v1/users", middlewares.CheckAuth))
	setupMatchesRoutes(engine.Group("v1/matches", middlewares.CheckAuth, middlewares.IsAdmin))
}

func setupUnauthenticatedRoutes(authGroup *gin.RouterGroup) {
	authGroup.POST("/sign-in", login)
}

func setupAuthRoutes(authGroup *gin.RouterGroup) {
	authGroup.POST("/", middlewares.IsAdmin, registerUser)
	authGroup.PUT("/", updateUser)
}

func setupMatchesRoutes(authGroup *gin.RouterGroup) {
	authGroup.POST("/:season", initializeMatches)
	authGroup.PUT("/", updateMatches)
}
