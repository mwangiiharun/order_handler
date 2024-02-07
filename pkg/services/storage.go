package services

import (
	"context"
	"fmt"
	"net/http"

	"cloud.google.com/go/firestore/apiv1/firestorepb"
	"github.com/order_handler/pkg/storage"
)

type FirestoreServiceConfig struct {
	ServiceConfig
	storage.FirestoreConfig
}

type FirestoreService struct {
	store  *storage.Firestore
	config *ServiceConfig
	Service
}

func (s *FirestoreService) Start(ctx context.Context) error {
	e := s.Service.Start(ctx)
	if e != nil {
		return e
	}

	credentials_file, exists := s.config.meta["credentials_file"]
	if !exists {
		return fmt.Errorf("credentials_file not found in config")
	}
	

	store, err := storage.NewFirestore(ctx, &storage.FirestoreConfig{
		Credentials_file: credentials_file,
		Project_id:       s.config.meta["project_id"],
		Database_id:      s.config.meta["database_id"],
	})
	if err != nil {
		return err
	}
	s.store = store
	err = s.store.Connect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *FirestoreService) Fetch(r *http.Request, ctx context.Context) ([]interface{}, error) {
	// res, err := s.store.List(&firestorepb.ListDocumentsRequest{
	// 	Parent:       s.config.,
	// 	CollectionId: s.config.CollectionId,
	// }, ctx)
	// if err != nil {
	// 	return nil, err
	// }
	// return res, nil
	return nil, nil
}

func (s *FirestoreService) Get(r *http.Request, ctx context.Context) (interface{}, error) {
	res, err := s.store.Get(&firestorepb.GetDocumentRequest{}, ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *FirestoreService) Create(r *http.Request, ctx context.Context) (interface{}, error) {
	// var fields map[string]*firestorepb.Value
	// fields["name"] = &firestorepb.Value{ValueType: &firestorepb.Value_StringValue{StringValue: "test"}}

	// document := &firestorepb.Document{}

	// res, err := s.store.Create(&firestorepb.CreateDocumentRequest{
	// 	Parent:       s.config.Parent,
	// 	CollectionId: s.config.CollectionId,
	// 	Document:     document,
	// }, ctx)
	// if err != nil {
	// 	return nil, err
	// }
	// return res, nil
	return nil, nil
}

func (s *FirestoreService) Update(r *http.Request, tx context.Context) (interface{}, error) {
	res, err := s.store.Update(&firestorepb.UpdateDocumentRequest{}, tx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *FirestoreService) Delete(r *http.Request, tx context.Context) error {
	err := s.store.Delete(&firestorepb.DeleteDocumentRequest{}, tx)
	if err != nil {
		return err
	}
	return nil
}

func NewFirestore(ctx context.Context, config *ServiceConfig) *FirestoreService {
	return &FirestoreService{
		config: config,
	}
}
