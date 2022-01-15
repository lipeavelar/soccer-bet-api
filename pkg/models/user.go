package models

import (
	"time"

	"github.com/lipeavelar/soccer-bet-api/pkg/helpers"
	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"-"`
	PlainPassword  string `json:"password" gorm:"-"`
	IsAdmin        bool   `json:"isAdmin"`
	ChangePassword bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.PlainPassword), 10)
	if err != nil {
		return helpers.GenerateError(err)
	}
	u.Password = string(hashedPassword)
	return nil
}
