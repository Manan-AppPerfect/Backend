package repository

import "context"

type Repository[T any] interface {
	Create(ctx context.Context, entity *T) error
	GetByID(ctx context.Context, id int64) (*T, error)
	Update(ctx context.Context, id int64, entity *T) error
	Delete(ctx context.Context, id int64) error
}
