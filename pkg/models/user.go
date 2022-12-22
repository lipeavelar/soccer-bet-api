package models

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User model
type User struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"-" gorm:"column:password"`
	PlainPassword  string `json:"password" gorm:"-"`
	IsAdmin        bool   `json:"isAdmin"`
	ChangePassword bool   `json:"changePassword"`
	WantEmail      bool   `json:"wantEmail"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.PlainPassword == "" {
		return errors.New("user password is required")
	}

	hashedPassword, err := HashPassword(u.PlainPassword)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	return nil
}

func HashPassword(plainPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 10)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
