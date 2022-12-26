package initializers

import (
	"github.com/lipeavelar/soccer-bet-api/internal/repositories/betsrepo"
	"github.com/lipeavelar/soccer-bet-api/internal/repositories/matchesrepo"
	"github.com/lipeavelar/soccer-bet-api/internal/services/betssrv"
)

func BetsService() (betssrv.BetsService, error) {
	betsRepo, err := betsrepo.NewBetsRepo()
	if err != nil {
		return nil, err
	}
	matchesRepo, err := matchesrepo.NewMatchesRepo()
	if err != nil {
		return nil, err
	}
	return betssrv.NewBetsService(betssrv.BetRepositories{
		Matches: matchesRepo,
		Bets:    betsRepo,
	}), nil
}
