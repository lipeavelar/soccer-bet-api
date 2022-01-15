package auth

import (
	authrepo "github.com/lipeavelar/soccer-bet-api/internal/repositories/auth"
	"github.com/lipeavelar/soccer-bet-api/pkg/models"
)

type authService struct {
	repository authrepo.UserRepo
}

// AuthService is the service for authentication
type AuthService interface {
	CreateUser(user models.User) (models.User, error)
}

// NewAuthService returns a new AuthService
func NewAuthService(repo authrepo.UserRepo) AuthService {
	return &authService{
		repository: repo,
	}
}
