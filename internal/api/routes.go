package api

import "github.com/gin-gonic/gin"

// SetupRoutes sets up the routes for the API
func SetupRoutes(engine *gin.Engine) {
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
