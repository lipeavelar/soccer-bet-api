package matchessrv

import (
	"encoding/json"
	"time"

	"github.com/lipeavelar/soccer-bet-api/internal/repositories/matchesrepo"
	"github.com/lipeavelar/soccer-bet-api/pkg/helpers"
	"github.com/lipeavelar/soccer-bet-api/pkg/models"
)

type matchesService struct {
	repository matchesrepo.MatchesRepo
}

// MatchesService is the service for matches
type MatchesService interface {
	InitializeMatches(season int) error
	UpdateMatches() error
	getMatchesFromAPI() ([]models.MatchResponse, error)
}

// NewMatchesService returns a new MatchesService
func NewMatchesService(repo matchesrepo.MatchesRepo) MatchesService {
	return &matchesService{
		repository: repo,
	}
}

func (srv *matchesService) getMatchesFromAPI() ([]models.MatchResponse, error) {
	matchesJSON, err := helpers.FootballAPIRequest("matches")
	if err != nil {
		return nil, err
	}
	var matchesRes models.MatchesResponse
	if err := json.Unmarshal(matchesJSON, &matchesRes); err != nil {
		return nil, err
	}

	return matchesRes.Matches, nil
}

func getMatchLocalDate(date time.Time) (time.Time, error) {
	timezone, err := helpers.GetTimezoneString("brazil")
	if err != nil {
		return time.Time{}, err
	}
	return helpers.ConvertToTimezone(date, timezone)
}
