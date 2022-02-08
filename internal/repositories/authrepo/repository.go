package authrepo

import (
	"github.com/lipeavelar/soccer-bet-api/database"
	"github.com/lipeavelar/soccer-bet-api/pkg/models"
	"gorm.io/gorm"
)

type userRepository struct {
	connection *gorm.DB
}

type UserRepo interface {
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User, userUpdates map[string]interface{}) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserByID(id string) (models.User, error)
}

// NewUsersRepo Returns an User repo object
func NewUsersRepo() (UserRepo, error) {
	conn, err := database.GetConnection()
	if err != nil {
		return &userRepository{}, err
	}
	return &userRepository{
		connection: conn,
	}, nil
}
