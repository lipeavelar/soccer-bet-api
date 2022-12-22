package models

import "time"

type Bet struct {
	ID            int `json:"id"`
	MatchID       int `json:"matchId"`
	UserID        int
	HomeTeamScore int `json:"homeTeamScore"`
	AwayTeamScore int `json:"awayTeamScore"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
