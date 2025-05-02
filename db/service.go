package db

import (
	"context"
	"gorm.io/gorm"
)

// Service defines business logic operations with optional queries.
type Service[T any] interface {
	Create(ctx context.Context, model *T) error
	Get(ctx context.Context, id any, queryFns ...func(*gorm.DB) *gorm.DB) (*T, error)
	Update(ctx context.Context, model *T) error
	Delete(ctx context.Context, id any) error
	List(ctx context.Context, queryFns ...func(*gorm.DB) *gorm.DB) ([]T, error)
}

// service implements the Service interface.
type service[T any] struct {
	repo Repository[T]
}

// NewService returns a new generic service using the provided repository.
func NewService[T any](repo Repository[T]) Service[T] {
	return &service[T]{repo}
}

// Create creates a new model.
func (s *service[T]) Create(ctx context.Context, model *T) error {
	return s.repo.Create(ctx, model)
}

// Get retrieves a model by ID with optional queries.
func (s *service[T]) Get(ctx context.Context, id any, queryFns ...func(*gorm.DB) *gorm.DB) (*T, error) {
	return s.repo.GetByID(ctx, id, queryFns...)
}

// Update updates an existing model.
func (s *service[T]) Update(ctx context.Context, model *T) error {
	return s.repo.Update(ctx, model)
}

// Delete removes a model by ID.
func (s *service[T]) Delete(ctx context.Context, id any) error {
	return s.repo.Delete(ctx, id)
}

// List lists all models with optional queries.
func (s *service[T]) List(ctx context.Context, queryFns ...func(*gorm.DB) *gorm.DB) ([]T, error) {
	return s.repo.List(ctx, queryFns...)
}
