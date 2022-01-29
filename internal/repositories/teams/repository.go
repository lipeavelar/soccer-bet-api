package teams

import (
	"github.com/lipeavelar/soccer-bet-api/database"
	"github.com/lipeavelar/soccer-bet-api/pkg/models"
	"gorm.io/gorm"
)

type teamsRepository struct {
	connection *gorm.DB
}

type TeamsRepo interface {
	GetTeamsName() ([]string, error)
	CreateTeams(teams []models.Team) error
}

func NewTeamsRepo() (TeamsRepo, error) {
	conn, err := database.GetConnection()
	if err != nil {
		return &teamsRepository{}, err
	}
	return &teamsRepository{
		connection: conn,
	}, nil
}
