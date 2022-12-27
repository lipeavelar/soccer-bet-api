package betsrepo

import "github.com/lipeavelar/soccer-bet-api/pkg/models"

func (repo *betsRepository) CreateBet(bet models.Bet) (models.Bet, error) {
	results := repo.connection.Create(&bet)
	return bet, results.Error
}

func (repo *betsRepository) UpdateBet(bet models.Bet) (models.Bet, error) {
	results := repo.connection.Model(&bet).Select("HomeTeamScore", "AwayTeamScore").Updates(&bet)
	return bet, results.Error
}

func (repo *betsRepository) GetBet(id int) (models.Bet, error) {
	var bet models.Bet
	results := repo.connection.First(&bet, id)
	return bet, results.Error
}

func (repo *betsRepository) GetBets(filters models.Bet) ([]models.Bet, error) {
	var bets []models.Bet
	results := repo.connection.Where(filters).Find(&bets)
	return bets, results.Error
}
