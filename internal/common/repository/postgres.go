package repository

import (
	"context"

	"gorm.io/gorm"
)



type PostgresRepo[T any] struct {
	db *gorm.DB
}

func NewPostgresRepository[T any] (db *gorm.DB) Repository[T] {
	return &PostgresRepo[T]{
		db: db,
	}
}

//Implementation

func(r *PostgresRepo[T]) Create(ctx context.Context, entity *T) error {
	return r.db.WithContext(ctx).Create(entity).Error
}

func(r *PostgresRepo[T]) GetByID(ctx context.Context, id int64) (*T, error) {
	var entity T

	err := r.db.WithContext(ctx).
		First(&entity, id).
		Error

	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func(r *PostgresRepo[T]) Update(ctx context.Context, id int64, entity *T) error {

	return r.db.WithContext(ctx).
		Model(entity).
		Where("id = ?", id).
		Updates(entity).
		Error

}

func(r *PostgresRepo[T]) Delete(ctx context.Context, id int64) error {
	var entity T

	return r.db.WithContext(ctx).
		Delete(&entity, id).
		Error
		
}