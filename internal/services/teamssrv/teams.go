package teamssrv

import "github.com/lipeavelar/soccer-bet-api/internal/repositories/teamsrepo"

type teamsResponse struct {
	Teams []teamResponse `json:"teams"`
}

type teamResponse struct {
	Name     string `json:"name"`
	Alias    string `json:"shortName"`
	Acronym  string `json:"tla"`
	CrestURL string `json:"crestUrl"`
}

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
