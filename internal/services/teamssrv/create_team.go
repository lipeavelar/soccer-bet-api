package teamssrv

import (
	"encoding/json"

	"github.com/lipeavelar/soccer-bet-api/pkg/helpers"
	"github.com/lipeavelar/soccer-bet-api/pkg/models"
)

func (t *teamsService) CreateTeams() error {
	teamsJSON, err := helpers.HttpRequest("teams")
	if err != nil {
		return err
	}
	var teamRes teamsResponse
	if err := json.Unmarshal(teamsJSON, &teamRes); err != nil {
		return err
	}

	existingTeams, err := t.repository.GetTeamsName()
	if err != nil {
		return err
	}

	teamsToCreate := make([]models.Team, 0)
	for _, team := range teamRes.Teams {
		if !helpers.ContainsString(existingTeams, team.Name) {
			teamToCreate := models.Team{
				Name:     team.Name,
				Alias:    team.Alias,
				Acronym:  team.Acronym,
				CrestURL: team.CrestURL,
			}
			teamsToCreate = append(teamsToCreate, teamToCreate)
		}
	}
	if err := t.repository.CreateTeams(teamsToCreate); len(teamsToCreate) > 0 && err != nil {
		return err
	}
	return nil
}
