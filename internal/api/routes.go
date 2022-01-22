package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lipeavelar/soccer-bet-api/internal/middlewares"
)

// SetupRoutes sets up the routes for the API
func SetupRoutes(engine *gin.Engine) {
	setupUnauthenticatedRoutes(engine.Group("v1"))
	setupAuthRoutes(engine.Group("v1/auth", middlewares.CheckAuth))
}

func setupUnauthenticatedRoutes(authGroup *gin.RouterGroup) {
	authGroup.POST("/signin", login)
}

func setupAuthRoutes(authGroup *gin.RouterGroup) {
	authGroup.POST("/register", registerUser)
}
