package initializers

import (
	"github.com/lipeavelar/soccer-bet-api/internal/repositories/teamsrepo"
	"github.com/lipeavelar/soccer-bet-api/internal/services/teamssrv"
)

func TeamsService() (teamssrv.TeamsService, error) {
	teamsRepo, err := teamsrepo.NewTeamsRepo()
	if err != nil {
		return nil, err
	}
	return teamssrv.NewTeamsService(teamsRepo), nil
}
