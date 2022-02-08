package authsrv

import (
	"github.com/lipeavelar/soccer-bet-api/pkg/models"
)

func (srv *authService) CreateUser(user models.User) (models.User, error) {
	user.ChangePassword = true
	user.WantEmail = true
	registeredUser, err := srv.repository.CreateUser(user)
	if err != nil {
		return models.User{}, err
	}
	return registeredUser, nil
}
