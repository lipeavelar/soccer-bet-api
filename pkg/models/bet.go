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

type BetCreateRequest struct {
	MatchID       *int `json:"matchId" binding:"required,number,gte=0"`
	UserID        int
	HomeTeamScore *int `json:"homeTeamScore" binding:"required,number,gte=0"`
	AwayTeamScore *int `json:"awayTeamScore" binding:"required,number,gte=0"`
}

type BetUpdateRequest struct {
	HomeTeamScore *int `json:"homeTeamScore" binding:"required,number,gte=0"`
	AwayTeamScore *int `json:"awayTeamScore" binding:"required,number,gte=0"`
}
