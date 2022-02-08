package teamssrv

import (
	"github.com/lipeavelar/soccer-bet-api/pkg/helpers"
	"github.com/lipeavelar/soccer-bet-api/pkg/models"
)

func (t *teamsService) CreateTeams(teams []string) error {
	existingTeams, err := t.repository.GetTeamsName()
	if err != nil {
		return err
	}
	teamsToCreate := make([]models.Team, 0)
	for _, team := range teams {
		if !helpers.ContainsString(existingTeams, team) {
			teamToCreate := models.Team{
				Name: team,
			}
			teamsToCreate = append(teamsToCreate, teamToCreate)
		}
	}
	return t.repository.CreateTeams(teamsToCreate)
}
