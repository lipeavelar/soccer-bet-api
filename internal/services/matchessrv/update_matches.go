package matchessrv

import (
	"github.com/lipeavelar/soccer-bet-api/pkg/helpers"
	"github.com/lipeavelar/soccer-bet-api/pkg/models"
)

func (srv *matchesService) UpdateMatches() ([]models.Match, error) {
	currentSeason, err := srv.repository.GetCurrentSeason()
	if err != nil {
		return []models.Match{}, err
	}
	matchesRes, err := srv.getMatchesFromAPI(currentSeason)
	if err != nil {
		return []models.Match{}, err
	}

	matches, err := srv.repository.GetMatches(models.Match{
		Season: currentSeason,
	})
	if err != nil {
		return []models.Match{}, err
	}
	matchesToUpdate := make([]models.Match, 0)
	for _, match := range matches {
		newMatch, err := compareMatch(match, matchesRes)
		if err != nil {
			return []models.Match{}, err
		}
		if newMatch.ID > 0 {
			matchesToUpdate = append(matchesToUpdate, newMatch)
		}
	}

	return srv.repository.SaveMatches(matchesToUpdate)
}

func compareMatch(match models.Match, newMatches []models.MatchResponse) (models.Match, error) {
	for _, newMatchRes := range newMatches {
		newMatchDate, err := getMatchLocalDate(newMatchRes.Date)
		if err != nil {
			return models.Match{}, err
		}
		newMatchRes.Date = newMatchDate

		if match.APIMatchID == newMatchRes.ID && (helpers.CompareDates(newMatchRes.Date, match.MatchDate) != 0 ||
			newMatchRes.Score.FullTime.HomeTeam != match.HomeTeamScore ||
			newMatchRes.Score.FullTime.AwayTeam != match.AwayTeamScore) {
			return models.Match{
				ID:            match.ID,
				APIMatchID:    match.APIMatchID,
				MatchDate:     newMatchRes.Date,
				MatchDay:      match.MatchDay,
				HomeTeam:      match.HomeTeam,
				AwayTeam:      match.AwayTeam,
				Season:        match.Season,
				HomeTeamScore: newMatchRes.Score.FullTime.HomeTeam,
				AwayTeamScore: newMatchRes.Score.FullTime.AwayTeam,
			}, nil
		}
	}
	return models.Match{
		ID: 0,
	}, nil
}
