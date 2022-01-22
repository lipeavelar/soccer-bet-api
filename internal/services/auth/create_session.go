package auth

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/lipeavelar/soccer-bet-api/pkg/constants"
	"github.com/lipeavelar/soccer-bet-api/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func (srv *authService) CreateSession(user models.User) (string, error) {
	checkUser, err := srv.repository.GetUserByEmail(user.Email)
	if err != nil {
		return "", err
	}
	if checkUser.ID == 0 {
		return "", errors.New("wrong user/password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(user.Password)); err != nil {
		return "", errors.New("wrong user/password")
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(checkUser.ID),
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
	})

	secret := os.Getenv(constants.SecretTokenEnvKey)

	token, err := claims.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}
