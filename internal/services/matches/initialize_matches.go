package matches

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/lipeavelar/soccer-bet-api/pkg/models"
)

func (srv *matchesService) InitializeMatches(currentSeason int) error {
	apiUrl := os.Getenv("INFO_API_URL")
	apiToken := os.Getenv("INFO_API_TOKEN")
	apiMethod := os.Getenv("INFO_API_METHOD")
	client := &http.Client{}
	request, err := http.NewRequest(apiMethod, apiUrl, nil)
	if err != nil {
		return err
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Auth-Token", apiToken)
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	var matchesRes matchesResponse
	if err := json.Unmarshal(bodyBytes, &matchesRes); err != nil {
		return err
	}

	matches := make([]models.Match, len(matchesRes.Matches))
	for i, matchRes := range matchesRes.Matches {
		matches[i] = models.Match{
			APIMatchID:    matchRes.ID,
			MatchDate:     matchRes.Date,
			MatchDay:      matchRes.MatchDay,
			HomeTeam:      matchRes.HomeTeamName.Name,
			AwayTeam:      matchRes.AwayTeamName.Name,
			Season:        currentSeason,
			HomeTeamScore: -1,
			AwayTeamScore: -1,
		}
	}

	return srv.repository.CreateMatches(matches)
}
