package main

import (
	"os"

	"github.com/lipeavelar/soccer-bet-api/internal/api"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("API_PORT")
	engine := gin.Default()
	api.SetupRoutes(engine)
	engine.Run(":" + port)
}
