package api

import "github.com/gin-gonic/gin"

// SetupRoutes sets up the routes for the API
func SetupRoutes(engine *gin.Engine) {
	setupAuthRoutes(engine.Group("v1/auth"))
}

func setupAuthRoutes(authGroup *gin.RouterGroup) {
	authGroup.POST("/register", register)
}
