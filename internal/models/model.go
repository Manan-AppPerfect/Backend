package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// model.go defines the bluePrint for the data

type Team struct {
	ID 			int64	`gorm:"primaryKey"`
	Name 		string	`gorm:"size:100;not null;unique"`
	Age 		int
	PlayerCount int
	Points 		int		`gorm:"default:0"`

	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	gorm.DeletedAt	`gorm:"index"`
}

func (t *Team) AddPlayerCount(p int) error {
	if p < 0 {
		return errors.New("playerCount cannot be negative")
	}
	t.PlayerCount += p
	return nil
}

func (t *Team) AddPoints(p int) error {
	if p < 0 {
		return errors.New("points cannot be negative")
	}
	t.Points += p
	return nil
}