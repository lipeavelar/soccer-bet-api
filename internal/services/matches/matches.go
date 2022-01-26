package matches

import (
	"time"

	"github.com/gin-gonic/gin"
	matchesrepo "github.com/lipeavelar/soccer-bet-api/internal/repositories/matches"
)

type matchesService struct {
	repository matchesrepo.MatchesRepo
	context    *gin.Context
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
}

// NewMatchesService returns a new MatchesService
func NewMatchesService(repo matchesrepo.MatchesRepo, c *gin.Context) MatchesService {
	return &matchesService{
		repository: repo,
		context:    c,
	}
}
