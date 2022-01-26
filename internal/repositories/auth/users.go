package auth

import (
	"github.com/lipeavelar/soccer-bet-api/pkg/models"
)

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

// CreateUser creates an user on database
func (repo *userRepository) CreateUser(user models.User) (models.User, error) {
	results := repo.connection.Create(&user)
	return user, results.Error
}

// UpdateUser updates user informations on database
func (repo *userRepository) UpdateUser(user models.User, userUpdates map[string]interface{}) (models.User, error) {
	results := repo.connection.Model(&user).Updates(userUpdates)
	return user, results.Error
}
