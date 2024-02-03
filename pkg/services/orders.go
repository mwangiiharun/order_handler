package services

import (
	"context"
	"net/http"

	"cloud.google.com/go/firestore/apiv1/firestorepb"
	"github.com/order_handler/pkg/storage"
)


type OrdersService struct {
	Name string
	Store *storage.Firestore
	Parent string
	CollectionId string
	Service
}

func (o *OrdersService) Start(ctx context.Context) error {
	return nil
}

func (o *OrdersService) Stop(ctx context.Context) error {
	return nil
}

func (o *OrdersService) Fetch(r *http.Request, tx context.Context) ([]interface{}, error) {
	res, err := o.Store.List(&firestorepb.ListDocumentsRequest{}, tx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (o *OrdersService) Get(r *http.Request, tx context.Context) (interface{}, error) {
	return nil, nil
}

func (o *OrdersService) Create(r *http.Request, tx context.Context) (interface{}, error) {

	return nil, nil
}

func (o *OrdersService) Update(r *http.Request, tx context.Context) (interface{}, error) {

	return nil, nil
}	

func (o *OrdersService) Delete(r *http.Request, tx context.Context) error {
	return nil
}

