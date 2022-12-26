package initializers

import (
	"github.com/lipeavelar/soccer-bet-api/internal/repositories/matchesrepo"
	"github.com/lipeavelar/soccer-bet-api/internal/services/matchessrv"
)

func MatchesService() (matchessrv.MatchesService, error) {
	repo, err := matchesrepo.NewMatchesRepo()
	if err != nil {
		return nil, err
	}
	return matchessrv.NewMatchesService(repo), nil
}
