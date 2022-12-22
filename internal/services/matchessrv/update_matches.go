package matchessrv

import (
	"github.com/lipeavelar/soccer-bet-api/pkg/helpers"
	"github.com/lipeavelar/soccer-bet-api/pkg/models"
)

func (srv *matchesService) UpdateMatches() error {
	matchesRes, err := srv.getMatchesFromAPI()
	if err != nil {
		return err
	}
	currentSeason, err := srv.repository.GetCurrentSeason()
	if err != nil {
		return err
	}

	matches, err := srv.repository.GetMatches(models.Match{
		Season: currentSeason,
	})
	if err != nil {
		return err
	}
	matchesToUpdate := make([]models.Match, 0)
	for _, match := range matches {
		newMatch := compareMatch(match, matchesRes)
		if newMatch.ID > 0 {
			newMatchDate, err := getMatchLocalDate(newMatch.MatchDate)
			if err != nil {
				return err
			}
			newMatch.MatchDate = newMatchDate
			matchesToUpdate = append(matchesToUpdate, newMatch)
		}
	}

	return srv.repository.SaveMatches(matchesToUpdate)
}

func compareMatch(match models.Match, newMatches []models.MatchResponse) models.Match {
	for _, newMatchRes := range newMatches {
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
			}
		}
	}
	return models.Match{
		ID: 0,
	}
}
