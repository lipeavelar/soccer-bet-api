package auth

import (
	"github.com/gin-gonic/gin"
	authrepo "github.com/lipeavelar/soccer-bet-api/internal/repositories/auth"
	"github.com/lipeavelar/soccer-bet-api/pkg/models"
)

type authService struct {
	repository authrepo.UserRepo
	context    *gin.Context
}

// AuthService is the service for authentication
type AuthService interface {
	CreateUser(user models.User) (models.User, error)
	CreateSession(user models.User) (string, error)
	UpdateUser(userUpdates map[string]interface{}) (models.User, error)
}

// NewAuthService returns a new AuthService
func NewAuthService(repo authrepo.UserRepo, c *gin.Context) AuthService {
	return &authService{
		repository: repo,
		context:    c,
	}
}
