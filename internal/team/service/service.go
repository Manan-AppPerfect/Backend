package service

// Service is the file where you provide Services like CRUD

import (
	"context"
	"errors"

	"github.com/Manan-AppPerfect/Backend/internal/common/repository"
	"github.com/Manan-AppPerfect/Backend/internal/gateway/slack"
	"github.com/Manan-AppPerfect/Backend/internal/gateway/user"
	"github.com/Manan-AppPerfect/Backend/internal/models"
)

type Service struct {
	repo 			repository.Repository[models.Team]
	userGateway 	user.Gateway
	slackGateway	slack.Gateway
}

func NewService(
	r 				repository.Repository[models.Team],
	userGateway 	user.Gateway,
	slackGateway	slack.Gateway,
) *Service {
	return &Service{
		repo: r,
		userGateway: userGateway,
		slackGateway: slackGateway,
	}
}

func (s *Service) CreateTeam(ctx context.Context, name string, age, playerCount, points int) (*models.Team, error){

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
		Name: name,
		Age: age,
		PlayerCount: playerCount,
		Points: points,
	}

	if err := s.repo.Create(ctx, team); err != nil{
		return nil, err
	}

	message := "New Team Created: " + team.Name

	err = s.slackGateway.SendMessage(ctx, message)
	if err != nil {
		return nil, err 
	}

	return team, nil
}

func (s *Service) GetTeamByID(ctx context.Context, id int64) (*models.Team, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) AddPlayers(ctx context.Context, id int64, count int) error {
	team, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if err := team.AddPlayerCount(count); err != nil {
		return err
	}

	return s.repo.Update(ctx, id, team)
}

func (s *Service) AddPoints(ctx context.Context, id int64, count int) error {
	team, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if err := team.AddPoints(count); err != nil {
		return err
	}

	return s.repo.Update(ctx, id, team)
}