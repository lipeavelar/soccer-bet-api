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
		context.JSON(http.StatusBadRequest, helpers.GenerateError("invalid season value", err))
		return
	}

	// Initialize matches season
	matchService, err := initializers.MatchesService()
	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError("internal server error", err))
		return
	}
	if err := matchService.InitializeMatches(currentSeason); err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError("internal server error", err))
		return
	}

	teamsService, err := initializers.TeamsService()
	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError("internal server error", err))
		return
	}
	if err := teamsService.CreateTeams(); err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError("internal server error", err))
		return
	}
	context.JSON(http.StatusOK, map[string]string{
		"result": "Matches initialized with success",
	})
}

func updateMatches(context *gin.Context) {
	matchService, err := initializers.MatchesService()
	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError("internal server error", err))
		return
	}
	if err := matchService.UpdateMatches(); err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError("internal server error", err))
		return
	}
	context.JSON(http.StatusOK, map[string]string{
		"result": "Matches updated with success",
	})
}
