package auth

import (
	"github.com/lipeavelar/soccer-bet-api/pkg/models"
	"gorm.io/gorm"
)

type userRepository struct {
	connection *gorm.DB
}

type UserRepo interface {
	CreateUser(user models.User) (models.User, error)
}
