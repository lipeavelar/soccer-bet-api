package models

type ScoreResponse struct {
	FullTime struct {
		HomeTeam int `json:"homeTeam"`
		AwayTeam int `json:"awayTeam"`
	} `json:"fullTime"`
}
