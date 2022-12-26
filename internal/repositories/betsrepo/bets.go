package betsrepo

import "github.com/lipeavelar/soccer-bet-api/pkg/models"

func (repo *betsRepository) CreateBet(bet models.Bet) error {
	results := repo.connection.Create(&bet)
	return results.Error
}

func (repo *betsRepository) UpdateBet(bet models.Bet) error {
	results := repo.connection.Model(&bet).Select("HomeTeamScore", "AwayTeamScore").Updates(&bet)
	return results.Error
}

func (repo *betsRepository) GetBet(id int) (models.Bet, error) {
	var bet models.Bet
	results := repo.connection.First(&bet, id)
	return bet, results.Error
}
