package betssrv

import (
	"github.com/lipeavelar/soccer-bet-api/pkg/models"
)

func (srv *betsService) CreateBet(bet models.Bet) (models.Bet, error) {
	if err := srv.checkMatchIsBetable(bet.MatchID); err != nil {
		return models.Bet{}, err
	}
	err := srv.repositories.Bets.CreateBet(bet)
	if err != nil {
		return models.Bet{}, err
	}
	return bet, nil
}
