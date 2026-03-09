package repository

import (
	"context"

	"github.com/Manan-AppPerfect/Backend/internal/models"
)

//go:generate mockgen -destination=./mocks/mock_repository.go -package=mocks github.com/Manan-AppPerfect/Backend/internal/common/repository TeamRepository

type Repository[T any] interface {
	Create(ctx context.Context, entity *T) error
	GetByID(ctx context.Context, id int64) (*T, error)
	Update(ctx context.Context, id int64, entity *T) error
	Delete(ctx context.Context, id int64) error
}


type TeamRepository interface {
	Repository[models.Team]
}