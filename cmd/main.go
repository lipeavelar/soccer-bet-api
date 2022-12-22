package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/lipeavelar/soccer-bet-api/internal/api"
	"github.com/lipeavelar/soccer-bet-api/internal/jobs"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("API_PORT")
	engine := gin.Default()
	engine.SetTrustedProxies(nil)
	api.SetupRoutes(engine)
	engine.Run(":" + port)

	//TODO: remove this from here and create a separate scheduler (maybe a scheduler is not even necessary, just schedule this to execute as cmd through command)
	scheduler := gocron.NewScheduler(time.UTC)
	_, err := scheduler.Every(1).Day().At("05:00").Do(jobs.UpdateMatches)
	if err != nil {
		// TODO: send email to admin
		fmt.Printf("======= Error scheduling job: %s\n\n", err)
	}
}
