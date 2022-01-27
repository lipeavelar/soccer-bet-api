package jobs

import (
	matchesrepo "github.com/lipeavelar/soccer-bet-api/internal/repositories/matches"
	matchessrv "github.com/lipeavelar/soccer-bet-api/internal/services/matches"
)

func UpdateMatches() error {
	repo, err := matchesrepo.NewMatchesRepo()
	if err != nil {
		return err
	}
	matchService := matchessrv.NewMatchesService(repo)
	if err := matchService.UpdateMatches(); err != nil {
		return err
	}
	return nil
}
