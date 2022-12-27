package matchessrv

import (
	"github.com/lipeavelar/soccer-bet-api/pkg/models"
)

func (srv *matchesService) InitializeMatches(currentSeason int) ([]models.Match, error) {
	matchesRes, err := srv.getMatchesFromAPI(currentSeason)
	if err != nil {
		return []models.Match{}, err
	}

	matches := make([]models.Match, len(matchesRes))

	for i, matchRes := range matchesRes {
		matchDay, err := getMatchLocalDate(matchRes.Date)
		if err != nil {
			return []models.Match{}, err
		}
		matches[i] = models.Match{
			APIMatchID:    matchRes.ID,
			MatchDate:     matchDay,
			MatchDay:      matchRes.MatchDay,
			HomeTeam:      matchRes.HomeTeamName.Name,
			AwayTeam:      matchRes.AwayTeamName.Name,
			Season:        currentSeason,
			HomeTeamScore: -1,
			AwayTeamScore: -1,
		}
	}

	return srv.repository.SaveMatches(matches)
}
