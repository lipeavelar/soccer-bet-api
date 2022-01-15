package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	authrepo "github.com/lipeavelar/soccer-bet-api/internal/repositories/auth"
	authsrv "github.com/lipeavelar/soccer-bet-api/internal/services/auth"
	"github.com/lipeavelar/soccer-bet-api/pkg/helpers"
	"github.com/lipeavelar/soccer-bet-api/pkg/models"
)

func register(context *gin.Context) {
	var user models.User

	if err := context.BindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, helpers.GenerateError(errors.New("invalid json")))
	}

	repo, err := authrepo.NewUserRepo()
	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError(errors.New("internal server error")))
	}
	userService := authsrv.NewAuthService(repo)
	createdUser, err := userService.CreateUser(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError(errors.New("internal server error")))
	}
	context.JSON(http.StatusOK, createdUser)
}
