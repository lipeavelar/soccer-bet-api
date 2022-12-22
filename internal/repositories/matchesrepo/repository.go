package matchesrepo

import (
	"github.com/lipeavelar/soccer-bet-api/database"
	"github.com/lipeavelar/soccer-bet-api/pkg/models"
	"gorm.io/gorm"
)

type matchesRepository struct {
	connection *gorm.DB
}

type MatchesRepo interface {
	SaveMatches(matches []models.Match) error
	GetMatches(filters models.Match) ([]models.Match, error)
	GetMatch(id int) (models.Match, error)
	GetCurrentSeason() (int, error)
}

// NewMatchesRepo Returns an Match repository object
func NewMatchesRepo() (MatchesRepo, error) {
	conn, err := database.GetConnection()
	if err != nil {
		return &matchesRepository{}, err
	}
	return &matchesRepository{
		connection: conn,
	}, nil
}
