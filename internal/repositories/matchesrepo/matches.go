package matchesrepo

import "github.com/lipeavelar/soccer-bet-api/pkg/models"

func (repo *matchesRepository) SaveMatches(matches []models.Match) error {
	results := repo.connection.Save(&matches)
	return results.Error
}

func (repo *matchesRepository) GetMatchesBySeason(season int) ([]models.Match, error) {
	var matches []models.Match
	results := repo.connection.Where("season = ?", season).Find(&matches)
	return matches, results.Error
}

func (repo *matchesRepository) GetCurrentSeason() (int, error) {
	var season int
	results := repo.connection.Table("matches").Select("max(season)").Scan(&season)
	return season, results.Error
}
