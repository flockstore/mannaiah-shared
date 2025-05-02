package db

import (
	"context"
	"gorm.io/gorm"
)

// Repository defines CRUD operations for any model T.
type Repository[T any] interface {
	Create(ctx context.Context, model *T) error
	GetByID(ctx context.Context, id any) (*T, error)
	Update(ctx context.Context, model *T) error
	Delete(ctx context.Context, id any) error
	List(ctx context.Context) ([]T, error)
	Async() AsyncRepository[T]
}

// AsyncRepository defines async versions of CRUD operations.
type AsyncRepository[T any] interface {
	CreateAsync(ctx context.Context, model *T) <-chan error
	GetByIDAsync(ctx context.Context, id any) <-chan Result[T]
	UpdateAsync(ctx context.Context, model *T) <-chan error
	DeleteAsync(ctx context.Context, id any) <-chan error
	ListAsync(ctx context.Context) <-chan Result[[]T]
}

// Result wraps data or error for async calls.
type Result[T any] struct {
	Data T
	Err  error
}

// repository is the default implementation of Repository.
type repository[T any] struct {
	db *gorm.DB
}

// NewRepository returns a new generic repository.
func NewRepository[T any](db *gorm.DB) Repository[T] {
	return &repository[T]{db}
}

// Create inserts a new record into the database.
func (r *repository[T]) Create(ctx context.Context, model *T) error {
	return r.db.WithContext(ctx).Create(model).Error
}

// GetByID retrieves a record by its primary key.
func (r *repository[T]) GetByID(ctx context.Context, id any) (*T, error) {
	var out T
	err := r.db.WithContext(ctx).First(&out, id).Error
	return &out, err
}

// Update saves the given model, updating fields by primary key.
func (r *repository[T]) Update(ctx context.Context, model *T) error {
	return r.db.WithContext(ctx).Save(model).Error
}

// Delete removes a record by its primary key.
func (r *repository[T]) Delete(ctx context.Context, id any) error {
	return r.db.WithContext(ctx).Delete(new(T), id).Error
}

// List retrieves all records of type T.
func (r *repository[T]) List(ctx context.Context) ([]T, error) {
	var out []T
	err := r.db.WithContext(ctx).Find(&out).Error
	return out, err
}

// Async returns an async wrapper for the repository.
func (r *repository[T]) Async() AsyncRepository[T] {
	return &asyncRepository[T]{repo: r}
}

// asyncRepository wraps sync repo with goroutine-based methods.
type asyncRepository[T any] struct {
	repo *repository[T]
}

// CreateAsync creates a record in a new goroutine.
func (a *asyncRepository[T]) CreateAsync(ctx context.Context, model *T) <-chan error {
	ch := make(chan error, 1)
	go func() { ch <- a.repo.Create(ctx, model) }()
	return ch
}

// GetByIDAsync retrieves a record by ID in a new goroutine.
func (a *asyncRepository[T]) GetByIDAsync(ctx context.Context, id any) <-chan Result[T] {
	ch := make(chan Result[T], 1)
	go func() {
		res, err := a.repo.GetByID(ctx, id)
		var out T
		if res != nil {
			out = *res
		}
		ch <- Result[T]{Data: out, Err: err}
	}()
	return ch
}

// UpdateAsync updates a record in a new goroutine.
func (a *asyncRepository[T]) UpdateAsync(ctx context.Context, model *T) <-chan error {
	ch := make(chan error, 1)
	go func() { ch <- a.repo.Update(ctx, model) }()
	return ch
}

// DeleteAsync deletes a record by ID in a new goroutine.
func (a *asyncRepository[T]) DeleteAsync(ctx context.Context, id any) <-chan error {
	ch := make(chan error, 1)
	go func() { ch <- a.repo.Delete(ctx, id) }()
	return ch
}

// ListAsync retrieves all records in a new goroutine.
func (a *asyncRepository[T]) ListAsync(ctx context.Context) <-chan Result[[]T] {
	ch := make(chan Result[[]T], 1)
	go func() {
		data, err := a.repo.List(ctx)
		ch <- Result[[]T]{Data: data, Err: err}
	}()
	return ch
}
