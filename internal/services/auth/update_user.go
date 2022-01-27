package auth

import (
	"github.com/lipeavelar/soccer-bet-api/pkg/models"
)

func (srv *authService) UpdateUser(userUpdates map[string]interface{}, loggedUser models.User) (models.User, error) {
	if _, ok := userUpdates["password"]; ok {
		var err error
		userUpdates["password"], err = models.HashPassword(userUpdates["password"].(string))
		if err != nil {
			return models.User{}, err
		}
	}
	updatedUser, err := srv.repository.UpdateUser(loggedUser, userUpdates)
	if err != nil {
		return models.User{}, err
	}
	return updatedUser, nil
}
