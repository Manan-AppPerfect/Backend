package service

// Service is the file where you provide Services like CRUD

import (
	"errors"

	"github.com/Manan-AppPerfect/Backend/internal/team/models"
	"github.com/Manan-AppPerfect/Backend/internal/team/repository"
)

type Service struct {
	repo *repository.Repo
}

func NewService(r *repository.Repo) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) CreateTeam(id, name string, age int) (*models.Team, error){

	if name == "" {
		return nil, errors.New("Team name required")
	}

	if age <= 50 {
		return nil, errors.New("Team must be an elite level")
	}

	team := &models.Team{
		ID: id,
		Name: name,
		Age: age,
	}

	s.repo.Save(team)
	
	return team, nil
}

func (s *Service) AddPlayers(id string, count int) error {
	team := s.repo.Get(id)
	if team == nil {
		return errors.New("Team not found")
	}
	
	return team.AddPlayerCount(count)
}

func (s *Service) AddPoints(id string, count int) error {
	team := s.repo.Get(id)
	if team == nil {
		return errors.New("Team not found")
	}
	
	return team.AddPoints(count)
}