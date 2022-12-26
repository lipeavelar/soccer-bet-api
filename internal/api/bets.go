package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/lipeavelar/soccer-bet-api/internal/api/initializers"
	"github.com/lipeavelar/soccer-bet-api/pkg/helpers"
	"github.com/lipeavelar/soccer-bet-api/pkg/models"
)

func createBet(context *gin.Context) {
	var betCreateModel models.BetCreateRequest
	if err := context.BindJSON(&betCreateModel); err != nil {
		context.JSON(http.StatusBadRequest, helpers.GenerateError("invalid bet json", err))
		return
	}
	bet := models.Bet{
		MatchID:       *betCreateModel.MatchID,
		HomeTeamScore: *betCreateModel.HomeTeamScore,
		AwayTeamScore: *betCreateModel.AwayTeamScore,
	}

	if userRaw, ok := context.Get("user"); !ok {
		context.JSON(http.StatusBadRequest, helpers.GenerateError("missing user information", errors.New("missing user information")))
		return
	} else {
		user, ok := userRaw.(models.User)
		if !ok {
			context.JSON(http.StatusBadRequest, helpers.GenerateError("invalid user", errors.New("invalid user")))
			return
		}
		bet.UserID = user.ID
	}

	betsService, err := initializers.BetsService()
	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError("internal server error", err))
		return
	}
	createdBet, err := betsService.CreateBet(bet)
	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError("internal server error", err))
		return
	}
	context.JSON(http.StatusOK, createdBet)
}

func updateBet(context *gin.Context) {
	betID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, helpers.GenerateError("invalid bet id", err))
		return
	}
	var bet models.BetUpdateRequest
	if err := context.BindJSON(&bet); err != nil {
		context.JSON(http.StatusBadRequest, helpers.GenerateError("invalid bet json", err))
		return
	}
	updateBet := models.Bet{
		ID:            betID,
		HomeTeamScore: *bet.HomeTeamScore,
		AwayTeamScore: *bet.AwayTeamScore,
	}

	betsService, err := initializers.BetsService()
	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError("internal server error", err))
		return
	}
	updatedBet, err := betsService.UpdateBet(updateBet)
	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError("internal server error", err))
		return
	}
	context.JSON(http.StatusOK, updatedBet)
}
