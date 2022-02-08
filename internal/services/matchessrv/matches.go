package matchessrv

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/lipeavelar/soccer-bet-api/internal/repositories/matchesrepo"
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
	GetTeamsBySeason(season int) ([]string, error)
	getMatchesFromAPI() ([]matchResponse, error)
}

// NewMatchesService returns a new MatchesService
func NewMatchesService(repo matchesrepo.MatchesRepo) MatchesService {
	return &matchesService{
		repository: repo,
	}
}

func (srv *matchesService) getMatchesFromAPI() ([]matchResponse, error) {
	apiUrl := os.Getenv("INFO_API_URL")
	apiToken := os.Getenv("INFO_API_TOKEN")
	apiMethod := os.Getenv("INFO_API_METHOD")
	client := &http.Client{}
	request, err := http.NewRequest(apiMethod, apiUrl, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Auth-Token", apiToken)
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var matchesRes matchesResponse
	if err := json.Unmarshal(bodyBytes, &matchesRes); err != nil {
		return nil, err
	}

	return matchesRes.Matches, nil
}
