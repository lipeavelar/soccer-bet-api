package middlewares

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	authrepo "github.com/lipeavelar/soccer-bet-api/internal/repositories/auth"

	"github.com/golang-jwt/jwt"
	"github.com/lipeavelar/soccer-bet-api/pkg/constants"
	"github.com/lipeavelar/soccer-bet-api/pkg/helpers"
)

func CheckAuth(context *gin.Context) {
	jwtToken := context.Request.Header.Get("Authorization")

	if jwtToken == "" || !strings.Contains(jwtToken, "Bearer") {
		context.AbortWithStatus(401)
		return
	}

	jwtToken = strings.Replace(jwtToken, "Bearer ", "", 1)
	secret := os.Getenv(constants.SecretTokenEnvKey)
	token, err := jwt.ParseWithClaims(jwtToken, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		context.AbortWithStatus(401)
		return
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		context.AbortWithStatus(401)
		return
	}

	repo, err := authrepo.NewUserRepo()
	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError(errors.New("internal server error")))
	}
	user, err := repo.GetUserByID(claims.Issuer)
	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.GenerateError(errors.New("internal server error")))
	}

	context.Set("user", user)
	context.Next()
}
