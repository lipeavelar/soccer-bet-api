package initializers

import (
	"github.com/lipeavelar/soccer-bet-api/internal/repositories/authrepo"
	"github.com/lipeavelar/soccer-bet-api/internal/services/authsrv"
)

func AuthService() (authsrv.AuthService, error) {
	repo, err := authrepo.NewUsersRepo()
	if err != nil {
		return nil, err
	}

	return authsrv.NewAuthService(repo), nil
}
