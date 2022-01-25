package auth

import (
	"errors"

	"github.com/lipeavelar/soccer-bet-api/pkg/models"
)

func (srv *authService) UpdateUser(userUpdates map[string]interface{}) (models.User, error) {
	currentUserRaw, _ := srv.context.Get("user")
	if currentUserRaw == nil {
		return models.User{}, errors.New("invalid user")
	}
	currentUser, ok := currentUserRaw.(models.User)
	if !ok {
		return models.User{}, errors.New("invalid user")
	}

	if _, ok := userUpdates["password"]; ok {
		var err error
		userUpdates["password"], err = models.HashPassword(userUpdates["password"].(string))
		if err != nil {
			return models.User{}, err
		}
	}
	updatedUser, err := srv.repository.UpdateUser(currentUser, userUpdates)
	if err != nil {
		return models.User{}, err
	}
	return updatedUser, nil
}
