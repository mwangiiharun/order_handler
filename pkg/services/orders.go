package services

import "context"

type OrdersService struct {
	storage *FirestoreService
	Service
}

func NewOrdersService(ctx context.Context, config *ServiceConfig) (*OrdersService, error) {
	return &OrdersService{
		Service: Service{
			config: config,
		},
	}, nil
}
