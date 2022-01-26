package matches

import "github.com/lipeavelar/soccer-bet-api/pkg/models"

func (repo *matchesRepository) CreateMatches(matches []models.Match) error {
	results := repo.connection.Create(&matches)
	return results.Error
}
