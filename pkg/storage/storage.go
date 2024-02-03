package storage

import "context"

type Storage interface {
	List(req interface{}, ctx context.Context) (interface{}, error)
	Get(req interface{}, ctx context.Context) (interface{}, error)
	Create(req interface{}, ctx context.Context) (interface{}, error)
	Update(req interface{}, ctx context.Context) (interface{}, error)
	Delete(req interface{}, ctx context.Context) error
	Connect(ctx context.Context) error
	Close() error
}
