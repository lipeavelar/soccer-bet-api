package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	matchesrepo "github.com/lipeavelar/soccer-bet-api/internal/repositories/matches"
	matchessrv "github.com/lipeavelar/soccer-bet-api/internal/services/matches"
	"github.com/lipeavelar/soccer-bet-api/pkg/helpers"
)

func initializeMatches(context *gin.Context) {
	currentSeason, err := strconv.Atoi(context.Param("season"))
	if err != nil {
		context.JSON(http.StatusBadRequest, helpers.GenerateError(errors.New("invalid season value")))
	}

	repo, err := matchesrepo.NewMatchesRepo()
	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError(errors.New("internal server error")))
	}
	matchService := matchessrv.NewMatchesService(repo)
	if err := matchService.InitializeMatches(currentSeason); err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError(errors.New("internal server error")))
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
