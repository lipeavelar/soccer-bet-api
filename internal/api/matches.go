package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lipeavelar/soccer-bet-api/internal/api/initializers"
	"github.com/lipeavelar/soccer-bet-api/pkg/helpers"
)

func initializeMatches(context *gin.Context) {
	currentSeason, err := strconv.Atoi(context.Param("season"))
	if err != nil {
		context.JSON(http.StatusBadRequest, helpers.GenerateError("invalid season value", err, gin.DefaultErrorWriter))
		return
	}

	// Initialize matches season
	matchService, err := initializers.MatchesService()
	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError("internal server error", err, gin.DefaultErrorWriter))
		return
	}
	matches, err := matchService.InitializeMatches(currentSeason)
	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError("internal server error", err, gin.DefaultErrorWriter))
		return
	}

	teamsService, err := initializers.TeamsService()
	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError("internal server error", err, gin.DefaultErrorWriter))
		return
	}
	if err := teamsService.CreateTeams(); err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError("internal server error", err, gin.DefaultErrorWriter))
		return
	}
	context.JSON(http.StatusOK, map[string]interface{}{
		"result":  "Matches initialized with success",
		"matches": matches,
	})
}

func updateMatches(context *gin.Context) {
	matchService, err := initializers.MatchesService()
	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError("internal server error", err, gin.DefaultErrorWriter))
		return
	}
	updatedMatches, err := matchService.UpdateMatches()
	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError("internal server error", err, gin.DefaultErrorWriter))
		return
	}
	context.JSON(http.StatusOK, map[string]interface{}{
		"result":  "Matches updated with success",
		"matches": updatedMatches,
	})
}
