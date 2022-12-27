package jobs

import (
	"github.com/lipeavelar/soccer-bet-api/internal/repositories/matchesrepo"
	"github.com/lipeavelar/soccer-bet-api/internal/services/matchessrv"
)

func UpdateMatches() error {
	repo, err := matchesrepo.NewMatchesRepo()
	if err != nil {
		return err
	}
	matchService := matchessrv.NewMatchesService(repo)
	if _, err := matchService.UpdateMatches(); err != nil {
		return err
	}
	return nil
}
