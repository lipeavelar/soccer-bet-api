package betsrepo

import (
	"github.com/lipeavelar/soccer-bet-api/database"
	"github.com/lipeavelar/soccer-bet-api/pkg/models"
	"gorm.io/gorm"
)

type betsRepository struct {
	connection *gorm.DB
}

type BetsRepo interface {
	CreateBet(bet models.Bet) (models.Bet, error)
	UpdateBet(bet models.Bet) (models.Bet, error)
	GetBet(betID int) (models.Bet, error)
	GetBets(filters models.Bet) ([]models.Bet, error)
}

// NewBetsRepo Returns an Match repository object
func NewBetsRepo() (BetsRepo, error) {
	conn, err := database.GetConnection()
	if err != nil {
		return &betsRepository{}, err
	}
	return &betsRepository{
		connection: conn,
	}, nil
}
