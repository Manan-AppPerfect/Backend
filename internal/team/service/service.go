package service

// Service is the file where you provide Services like CRUD

import (
	"context"
	"errors"

	"github.com/Manan-AppPerfect/Backend/internal/common/repository"
	"github.com/Manan-AppPerfect/Backend/internal/gateway/user"
	"github.com/Manan-AppPerfect/Backend/internal/team/models"
)

type Service struct {
	repo repository.Repository[models.Team]
	userGateway user.Gateway
}

func NewService(
	r repository.Repository[models.Team],
	userGateway user.Gateway,
) *Service {
	return &Service{
		repo: r,
		userGateway: userGateway,
	}
}

func (s *Service) CreateTeam(ctx context.Context, id int64, name string, age int) (*models.Team, error){

	user, err := s.userGateway.GetUser(ctx, 1)
	if err != nil {
		return nil, err
	}

	println("User from external service: ", user)

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

	if err := s.repo.Create(ctx, team); err != nil{
		return nil, err
	}
	
	return team, nil
}

func (s *Service) AddPlayers(ctx context.Context, id int64, count int) error {
	team, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	
	return team.AddPlayerCount(count)
}

func (s *Service) AddPoints(ctx context.Context, id int64, count int) error {
	team, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	
	return team.AddPoints(count)
}