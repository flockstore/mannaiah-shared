package endpoint

import (
	"context"

	"github.com/flockstore/mannaiah-shared/db"
	gk "github.com/go-kit/kit/endpoint"
)

// Endpoints defines CRUD operations for any model type.
type Endpoints[T any] struct {
	Create gk.Endpoint
	Get    gk.Endpoint
	Update gk.Endpoint
	Delete gk.Endpoint
	List   gk.Endpoint
}

// NewEndpoints builds endpoints for the given db.Service
func NewEndpoints[T any](svc db.Service[T]) Endpoints[T] {
	return Endpoints[T]{
		Create: makeCreateEndpoint(svc),
		Get:    makeGetEndpoint(svc),
		Update: makeUpdateEndpoint(svc),
		Delete: makeDeleteEndpoint(svc),
		List:   makeListEndpoint(svc),
	}
}

// makeCreateEndpoint makes a generic CRUD Create endpoint.
func makeCreateEndpoint[T any](svc db.Service[T]) gk.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest[T])
		err := svc.Create(ctx, req.Model)
		return Response[any]{Data: nil, Err: err}, nil
	}
}

// makeGetEndpoint makes a generic CRUD Get endpoint.
func makeGetEndpoint[T any](svc db.Service[T]) gk.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		data, err := svc.Get(ctx, req.ID)
		return Response[*T]{Data: data, Err: err}, nil
	}
}

// makeUpdateEndpoint makes a generic CRUD Update endpoint.
func makeUpdateEndpoint[T any](svc db.Service[T]) gk.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest[T])
		err := svc.Update(ctx, req.Model)
		return Response[any]{Data: nil, Err: err}, nil
	}
}

// makeDeleteEndpoint makes a generic CRUD Update endpoint.
func makeDeleteEndpoint[T any](svc db.Service[T]) gk.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		err := svc.Delete(ctx, req.ID)
		return Response[any]{Data: nil, Err: err}, nil
	}
}

// makeListEndpoint makes a generic CRUD Query (R) endpoint.
func makeListEndpoint[T any](svc db.Service[T]) gk.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListRequest)
		data, err := svc.List(ctx, req.QueryFns...)
		return Response[[]T]{Data: data, Err: err}, nil
	}
}
