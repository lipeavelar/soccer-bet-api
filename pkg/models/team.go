package models

import "time"

type Team struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" gorm:"column:team_name"`
	Alias     string    `json:"alias" gorm:"column:team_alias"`
	Acronym   string    `json:"acronym" gorm:"column:team_acronym"`
	CrestURL  string    `json:"crestURL" gorm:"column:team_crest_url"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
