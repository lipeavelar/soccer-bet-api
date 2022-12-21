package matchessrv

import (
	"encoding/json"
	"time"

	"github.com/lipeavelar/soccer-bet-api/internal/repositories/matchesrepo"
	"github.com/lipeavelar/soccer-bet-api/pkg/helpers"
)

type matchesService struct {
	repository matchesrepo.MatchesRepo
}

type matchesResponse struct {
	Matches []matchResponse `json:"matches"`
}

type matchResponse struct {
	ID           int           `json:"id"`
	Date         time.Time     `json:"utcDate"`
	MatchDay     int           `json:"matchday"`
	HomeTeamName teamResponse  `json:"homeTeam"`
	AwayTeamName teamResponse  `json:"awayTeam"`
	Score        scoreResponse `json:"score"`
}

type teamResponse struct {
	Name string `json:"name"`
}

type scoreResponse struct {
	FullTime struct {
		HomeTeam int `json:"homeTeam"`
		AwayTeam int `json:"awayTeam"`
	} `json:"fullTime"`
}

// MatchesService is the service for matches
type MatchesService interface {
	InitializeMatches(season int) error
	UpdateMatches() error
	getMatchesFromAPI() ([]matchResponse, error)
}

// NewMatchesService returns a new MatchesService
func NewMatchesService(repo matchesrepo.MatchesRepo) MatchesService {
	return &matchesService{
		repository: repo,
	}
}

func (srv *matchesService) getMatchesFromAPI() ([]matchResponse, error) {
	matchesJSON, err := helpers.HttpRequest("matches")
	if err != nil {
		return nil, err
	}
	var matchesRes matchesResponse
	if err := json.Unmarshal(matchesJSON, &matchesRes); err != nil {
		return nil, err
	}

	return matchesRes.Matches, nil
}
