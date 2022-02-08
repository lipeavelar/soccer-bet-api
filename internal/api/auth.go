package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lipeavelar/soccer-bet-api/internal/repositories/authrepo"
	"github.com/lipeavelar/soccer-bet-api/internal/services/authsrv"
	"github.com/lipeavelar/soccer-bet-api/pkg/helpers"
	"github.com/lipeavelar/soccer-bet-api/pkg/models"
)

func registerUser(context *gin.Context) {
	var user models.User

	if err := context.BindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, helpers.GenerateError(errors.New("invalid json")))
		return
	}

	repo, err := authrepo.NewUsersRepo()
	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError(errors.New("internal server error")))
		return
	}
	userService := authsrv.NewAuthService(repo)
	createdUser, err := userService.CreateUser(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError(errors.New("internal server error")))
		return
	}
	context.JSON(http.StatusOK, createdUser)
}

func updateUser(context *gin.Context) {
	userJson, err := context.GetRawData()
	if err != nil {
		context.JSON(http.StatusBadRequest, helpers.GenerateError(errors.New("invalid json")))
		return
	}
	var user map[string]interface{}

	if err := json.Unmarshal(userJson, &user); err != nil {
		context.JSON(http.StatusBadRequest, helpers.GenerateError(errors.New("invalid json")))
		return
	}

	repo, err := authrepo.NewUsersRepo()
	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError(errors.New("internal server error")))
		return
	}

	loggedUserRaw, _ := context.Get("user")
	loggedUser, ok := loggedUserRaw.(models.User)
	if !ok {
		context.JSON(http.StatusBadRequest, helpers.GenerateError(errors.New("invalid user")))
		return
	}

	userService := authsrv.NewAuthService(repo)
	updatedUser, err := userService.UpdateUser(user, loggedUser)
	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError(errors.New("internal server error")))
		return
	}
	context.JSON(http.StatusOK, updatedUser)
}

func login(context *gin.Context) {
	var user models.User

	if err := context.BindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, helpers.GenerateError(errors.New("invalid json")))
		return
	}

	repo, err := authrepo.NewUsersRepo()
	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError(errors.New("internal server error")))
		return
	}
	userService := authsrv.NewAuthService(repo)
	token, err := userService.CreateSession(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError(errors.New("internal server error")))
		return
	}
	context.JSON(http.StatusOK, token)
}
