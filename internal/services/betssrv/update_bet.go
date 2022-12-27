package betssrv

import "github.com/lipeavelar/soccer-bet-api/pkg/models"

func (srv *betsService) UpdateBet(bet models.Bet) (models.Bet, error) {
	oldBet, err := srv.repositories.Bets.GetBet(bet.ID)
	if err != nil {
		return models.Bet{}, err
	}
	if err := srv.checkMatchIsBetable(oldBet.MatchID); err != nil {
		return models.Bet{}, err
	}
	updatedBet, err := srv.repositories.Bets.UpdateBet(bet)
	if err != nil {
		return models.Bet{}, err
	}
	return updatedBet, nil
}
