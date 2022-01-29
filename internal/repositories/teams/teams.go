package teams

import "github.com/lipeavelar/soccer-bet-api/pkg/models"

func (repo *teamsRepository) GetTeamsName() ([]string, error) {
	var teams []string
	results := repo.connection.Table("teams").Select("name")
	return teams, results.Error
}

func (repo *teamsRepository) CreateTeams(teams []models.Team) error {
	results := repo.connection.Table("teams").Create(&teams)
	return results.Error
}
