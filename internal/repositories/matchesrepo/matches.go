package matchesrepo

import "github.com/lipeavelar/soccer-bet-api/pkg/models"

func (repo *matchesRepository) SaveMatches(matches []models.Match) error {
	results := repo.connection.Save(&matches)
	return results.Error
}

func (repo *matchesRepository) GetMatch(id int) (models.Match, error) {
	var match models.Match
	results := repo.connection.First(&match, id)
	return match, results.Error
}

func (repo *matchesRepository) GetMatches(filters models.Match) ([]models.Match, error) {
	var matches []models.Match
	results := repo.connection.Where(filters).Find(&matches)
	return matches, results.Error
}

func (repo *matchesRepository) GetCurrentSeason() (int, error) {
	var season int
	results := repo.connection.Table("matches").Select("max(season)").Scan(&season)
	return season, results.Error
}
