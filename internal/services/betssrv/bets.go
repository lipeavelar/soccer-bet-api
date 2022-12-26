package betssrv

import (
	"errors"
	"time"

	"github.com/lipeavelar/soccer-bet-api/internal/repositories/betsrepo"
	"github.com/lipeavelar/soccer-bet-api/internal/repositories/matchesrepo"
	"github.com/lipeavelar/soccer-bet-api/pkg/helpers"
	"github.com/lipeavelar/soccer-bet-api/pkg/models"
)

type betsService struct {
	repositories BetRepositories
}

// BetsService is the service for bets
type BetsService interface {
	CreateBet(bet models.Bet) (models.Bet, error)
	UpdateBet(bet models.Bet) (models.Bet, error)
}

type BetRepositories struct {
	Matches matchesrepo.MatchesRepo
	Bets    betsrepo.BetsRepo
}

// NewAuthService returns a new AuthService
func NewBetsService(repos BetRepositories) BetsService {
	return &betsService{
		repositories: repos,
	}
}

func (srv *betsService) checkMatchIsBetable(matchID int) error {
	match, err := srv.repositories.Matches.GetMatch(matchID)
	if err != nil {
		return err
	}
	tz, err := helpers.GetTimezoneString("brazil")
	if err != nil {
		return err
	}
	currentDate, err := helpers.ConvertToTimezone(time.Now().UTC(), tz)
	if err != nil {
		return err
	}
	if helpers.CompareDates(currentDate, match.MatchDate) > -1 {
		return errors.New("date to bet on this match already passed")
	}
	return nil
}
