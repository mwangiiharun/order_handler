package services

import (
	"github.com/order_handler/pkg/storage"
)

type OrdersService struct {
	Name         string
	Store        *storage.Firestore
	Parent       string
	CollectionId string
	Service
}
