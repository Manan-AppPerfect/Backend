package repository

import (
	"context"
	"errors"
	"sync"
)

type MemoryRepo[T any] struct{
	data 	map[int64]*T
	mu 		sync.RWMutex
	nextID  int64
}

func NewMemoryRepository[T any]() Repository[T] {
	return &MemoryRepo[T]{
		data: make(map[int64]*T),
		nextID: 1,
	}
}

func (r *MemoryRepo[T]) Create(ctx context.Context, entity*T) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.data[r.nextID] = entity
	r.nextID++

	return nil
}

func (r *MemoryRepo[T]) GetByID(ctx context.Context, id int64) (*T, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	entity, ok := r.data[id]
	if !ok {
		return nil, errors.New("not found")
	}

	return entity, nil
}

func (r *MemoryRepo[T]) Update(ctx context.Context, id int64, entity*T) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.data[id]; !ok {
		return errors.New("not found")
	}

	r.data[id] = entity

	return nil
}

func (r *MemoryRepo[T]) Delete(ctx context.Context, id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.data[id]; !ok {
		return errors.New("not found")
	}

	delete(r.data, id)

	return nil
}
