package endpoint

import "gorm.io/gorm"

// CreateRequest defines a generic model request to be
// parsed between endpoints and service-created.
type CreateRequest[T any] struct {
	Model *T
}

// GetRequest defines a generic model find request
// to be made between endpoints.
type GetRequest struct {
	ID any
}

// UpdateRequest defines a generic model request to be
// parsed between endpoints and service-updated.
type UpdateRequest[T any] struct {
	Model *T
}

// DeleteRequest defines a generic model request
// to be made between endpoints.
type DeleteRequest struct {
	ID any
}

type ListRequest struct {
	QueryFns []func(*gorm.DB) *gorm.DB
}

// Response always returns model(s) and an error
type Response[T any] struct {
	Data T
	Err  error
}
