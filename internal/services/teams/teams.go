package teams

import teamsrepo "github.com/lipeavelar/soccer-bet-api/internal/repositories/teams"

type teamsService struct {
	repository teamsrepo.TeamsRepo
}

type TeamsService interface {
	CreateTeams(teams []string) error
}

func NewTeamsService(repo teamsrepo.TeamsRepo) TeamsService {
	return &teamsService{
		repository: repo,
	}
}
