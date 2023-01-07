package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/lipeavelar/soccer-bet-api/internal/api"
	"github.com/lipeavelar/soccer-bet-api/internal/jobs"

	"github.com/gin-gonic/gin"
)

func main() {
	logfile, err := os.OpenFile("logs/info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to create request log file:", err)
	}

	errlogfile, err := os.OpenFile("logs/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to create request log file:", err)
	}

	gin.DefaultWriter = io.MultiWriter(logfile, os.Stdout)
	gin.DefaultErrorWriter = io.MultiWriter(errlogfile, os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(gin.DefaultWriter)

	port := os.Getenv("API_PORT")
	engine := gin.Default()
	engine.SetTrustedProxies(nil)
	api.SetupRoutes(engine)
	engine.Run(":" + port)

	//TODO: remove this from here and create a separate scheduler (maybe a scheduler is not even necessary, just schedule this to execute as cmd through command)
	scheduler := gocron.NewScheduler(time.UTC)
	_, err = scheduler.Every(1).Day().At("05:00").Do(jobs.UpdateMatches)
	if err != nil {
		// TODO: send email to admin
		fmt.Printf("======= Error scheduling job: %s\n\n", err)
	}
}
