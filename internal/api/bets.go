package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/lipeavelar/soccer-bet-api/internal/api/initializers"
	"github.com/lipeavelar/soccer-bet-api/pkg/helpers"
	"github.com/lipeavelar/soccer-bet-api/pkg/models"
)

func createBet(context *gin.Context) {
	var bet models.Bet
	if err := context.BindJSON(&bet); err != nil {
		context.JSON(http.StatusBadRequest, helpers.GenerateError("invalid json", err))
		return
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
