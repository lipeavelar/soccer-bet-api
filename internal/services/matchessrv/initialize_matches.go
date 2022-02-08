package matchessrv

import (
	"github.com/lipeavelar/soccer-bet-api/pkg/models"
)

func (srv *matchesService) InitializeMatches(currentSeason int) error {
	matchesRes, err := srv.getMatchesFromAPI()
	if err != nil {
		return err
	}

	matches := make([]models.Match, len(matchesRes))
	for i, matchRes := range matchesRes {
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

	return srv.repository.SaveMatches(matches)
}
