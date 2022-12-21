package teamssrv

import "github.com/lipeavelar/soccer-bet-api/internal/repositories/teamsrepo"

type teamsService struct {
	repository teamsrepo.TeamsRepo
}

type TeamsService interface {
	CreateTeams() error
}

func NewTeamsService(repo teamsrepo.TeamsRepo) TeamsService {
	return &teamsService{
		repository: repo,
	}
}
