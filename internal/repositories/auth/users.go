package auth

import (
	"github.com/lipeavelar/soccer-bet-api/database"
	"github.com/lipeavelar/soccer-bet-api/pkg/models"
)

// NewUserRepo Returns an User repo object
func NewUserRepo() (UserRepo, error) {
	conn, err := database.GetConnection()
	if err != nil {
		return &userRepository{}, err
	}
	return &userRepository{
		connection: conn,
	}, nil
}

// CreateUser creates an user on database
func (repo *userRepository) CreateUser(user models.User) (models.User, error) {
	results := repo.connection.Create(&user)
	return user, results.Error
}

func (repo *userRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	results := repo.connection.Where("email = ?", email).FirstOrInit(&user)

	return user, results.Error
}

func (repo *userRepository) GetUserByID(id string) (models.User, error) {
	var user models.User
	results := repo.connection.Where("id = ?", id).First(&user)

	return user, results.Error
}
