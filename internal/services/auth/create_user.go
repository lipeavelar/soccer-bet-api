package auth

import (
	"errors"

	"github.com/lipeavelar/soccer-bet-api/pkg/models"
)

func (srv *authService) CreateUser(user models.User) (models.User, error) {
	currentUserRaw, _ := srv.context.Get("user")
	if currentUserRaw == nil {
		return models.User{}, errors.New("invalid user")
	}
	if currentUser, ok := currentUserRaw.(models.User); !ok || !currentUser.IsAdmin {
		return models.User{}, errors.New("invalid user")
	}

	registeredUser, err := srv.repository.CreateUser(user)
	if err != nil {
		return models.User{}, err
	}
	return registeredUser, nil
}
