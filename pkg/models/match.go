package models

import "time"

type Match struct {
	ID            int       `json:"id"`
	APIMatchID    int       `json:"-"`
	HomeTeam      string    `json:"homeTeam"`
	AwayTeam      string    `json:"awayTeam"`
	MatchDate     time.Time `json:"matchDate"`
	Season        int       `json:"season"`
	MatchDay      int       `json:"matchDay"`
	HomeTeamScore int       `json:"homeTeamScore"`
	AwayTeamScore int       `json:"awayTeamScore"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
}

type MatchesResponse struct {
	Matches []MatchResponse `json:"matches"`
}

type MatchResponse struct {
	ID           int           `json:"id"`
	Date         time.Time     `json:"utcDate"`
	MatchDay     int           `json:"matchday"`
	HomeTeamName TeamResponse  `json:"homeTeam"`
	AwayTeamName TeamResponse  `json:"awayTeam"`
	Score        ScoreResponse `json:"score"`
}
