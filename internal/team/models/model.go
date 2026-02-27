package models

import "errors"

// model.go defines the bluePrint for the data

type Team struct {
	ID 			int64
	Name 		string
	Age 		int
	PlayerCount int
	Points 		int
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