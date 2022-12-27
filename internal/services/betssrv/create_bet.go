package betssrv

import (
	"fmt"

	"github.com/lipeavelar/soccer-bet-api/pkg/models"
)

func (srv *betsService) CreateBet(bet models.Bet) (models.Bet, error) {
	if bets, err := srv.repositories.Bets.GetBets(models.Bet{
		MatchID: bet.MatchID,
		UserID:  bet.UserID,
	}); err == nil && len(bets) > 0 {
		return models.Bet{}, fmt.Errorf("bet for match %d already exists for user %d", bet.MatchID, bet.UserID)
	}
	if err := srv.checkMatchIsBetable(bet.MatchID); err != nil {
		return models.Bet{}, err
	}
	createdBet, err := srv.repositories.Bets.CreateBet(bet)
	if err != nil {
		return models.Bet{}, err
	}
	return createdBet, nil
}
