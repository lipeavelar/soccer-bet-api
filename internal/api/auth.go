package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lipeavelar/soccer-bet-api/internal/api/initializers"
	"github.com/lipeavelar/soccer-bet-api/pkg/helpers"
	"github.com/lipeavelar/soccer-bet-api/pkg/models"
)

func registerUser(context *gin.Context) {
	var user models.User

	if err := context.BindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, helpers.GenerateError("invalid json", err))
		return
	}

	userService, err := initializers.AuthService()
	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError("internal server error", err))
		return
	}
	createdUser, err := userService.CreateUser(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError("internal server error", err))
		return
	}
	context.JSON(http.StatusOK, createdUser)
}

func updateUser(context *gin.Context) {
	userJson, err := context.GetRawData()
	if err != nil {
		context.JSON(http.StatusBadRequest, helpers.GenerateError("invalid json", err))
		return
	}
	var user map[string]interface{}

	if err := json.Unmarshal(userJson, &user); err != nil {
		context.JSON(http.StatusBadRequest, helpers.GenerateError("invalid json", err))
		return
	}

	loggedUserRaw, _ := context.Get("user")
	loggedUser, ok := loggedUserRaw.(models.User)
	if !ok {
		context.JSON(http.StatusBadRequest, helpers.GenerateError("invalid user", err))
		return
	}

	userService, err := initializers.AuthService()
	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError("internal server error", err))
		return
	}

	updatedUser, err := userService.UpdateUser(user, loggedUser)
	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError("internal server error", err))
		return
	}
	context.JSON(http.StatusOK, updatedUser)
}

func login(context *gin.Context) {
	var user models.User

	if err := context.BindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, helpers.GenerateError("invalid json", err))
		return
	}

	userService, err := initializers.AuthService()
	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError("internal server error", err))
		return
	}
	token, err := userService.CreateSession(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError("internal server error", err))
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
