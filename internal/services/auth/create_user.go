package auth

import "github.com/lipeavelar/soccer-bet-api/pkg/models"

func (a *authService) CreateUser(user models.User) (models.User, error) {
	if err := user.HashPassword(); err != nil {
		return models.User{}, err
	}
	registeredUser, err := a.repository.CreateUser(user)
	if err != nil {
		return models.User{}, err
	}
	return registeredUser, nil
}
