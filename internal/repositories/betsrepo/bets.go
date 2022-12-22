package betsrepo

import "github.com/lipeavelar/soccer-bet-api/pkg/models"

func (repo *betsRepository) CreateBet(bet models.Bet) error {
	results := repo.connection.Create(&bet)
	return results.Error
}
