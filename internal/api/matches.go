package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lipeavelar/soccer-bet-api/internal/repositories/matchesrepo"
	"github.com/lipeavelar/soccer-bet-api/internal/repositories/teamsrepo"
	"github.com/lipeavelar/soccer-bet-api/internal/services/matchessrv"
	"github.com/lipeavelar/soccer-bet-api/internal/services/teamssrv"
	"github.com/lipeavelar/soccer-bet-api/pkg/helpers"
)

func initializeMatches(context *gin.Context) {
	currentSeason, err := strconv.Atoi(context.Param("season"))
	if err != nil {
		context.JSON(http.StatusBadRequest, helpers.GenerateError(errors.New("invalid season value")))
		return
	}

	// Initialize matches season
	matchesRepo, err := matchesrepo.NewMatchesRepo()
	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError(errors.New("internal server error")))
		return
	}
	matchService := matchessrv.NewMatchesService(matchesRepo)
	if err := matchService.InitializeMatches(currentSeason); err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError(errors.New("internal server error")))
		return
	}

	teamsRepo, err := teamsrepo.NewTeamsRepo()
	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError(errors.New("internal server error")))
		return
	}
	teamsService := teamssrv.NewTeamsService(teamsRepo)
	if err := teamsService.CreateTeams(); err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError(errors.New("internal server error")))
		return
	}
}

func updateMatches(context *gin.Context) {
	repo, err := matchesrepo.NewMatchesRepo()
	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError(errors.New("internal server error")))
	}
	matchService := matchessrv.NewMatchesService(repo)
	if err := matchService.UpdateMatches(); err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError(errors.New("internal server error")))
	}
}
