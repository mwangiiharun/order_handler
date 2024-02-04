package services

import (
	"context"
	"fmt"
	"net/http"

	"cloud.google.com/go/firestore/apiv1/firestorepb"
	"github.com/order_handler/pkg/storage"
)

type IService interface {
	SetName(name string) error
	setConfig(ctx context.Context) error
	Start(ctx context.Context) error
	Stop(ctx context.Context) error

	Fetch(r *http.Request, tx context.Context) ([]interface{}, error)
	Get(r *http.Request, tx context.Context) (interface{}, error)
	Create(r *http.Request, tx context.Context) (interface{}, error)
	Update(r *http.Request, tx context.Context) (interface{}, error)
	Delete(r *http.Request, tx context.Context) error
}

type ServiceType int

const (
	Storage ServiceType = iota + 1
	Http
)

func (s ServiceType) String() string {
	return [...]string{"Storage", "Http"}[s-1]
}

func (s ServiceType) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}

func (s ServiceType) UnmarshalText(text []byte) error {
	switch string(text) {
	case "Storage":
		s = Storage
	case "Http":
		s = Http
	default:
		return fmt.Errorf("invalid ServiceType: %s", text)
	}
	return nil
}

type ServiceConfig struct {
	Name         string      `yaml:"name"`
	Type         ServiceType `yaml:"type"`
	Parent       string      `yaml:"parent"`
	CollectionId string      `yaml:"collectionId"`
}

type ServiceConfigs *map[string]ServiceConfig

type Service struct {
	Name   string
	type_  ServiceType
	config *ServiceConfig
	store  *storage.Firestore
}

func (s *Service) setConfig(ctx context.Context) error {
	configs := ctx.Value("configs").(ServiceConfigs)
	config, ok := (*configs)[s.Name]
	if !ok {
		return fmt.Errorf("Service %s not found", s.Name)
	}
	s.config = &config
	return nil
}

func (s *Service) Start(ctx context.Context) error {
	err := s.setConfig(ctx)
	if err != nil {
		return err
	}
	s.type_ = s.config.Type

	if ctx.Value("storage") != nil {
		s.store = ctx.Value("storage").(*storage.Firestore)
		return nil
	}

	s.store, err = storage.NewFirestore(ctx)
	if err != nil {
		return err
	}
	err = s.store.Connect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Stop(ctx context.Context) error {
	return s.store.Close()
}

func (s *Service) Fetch(r *http.Request, ctx context.Context) ([]interface{}, error) {
	res, err := s.store.List(&firestorepb.ListDocumentsRequest{
		Parent:       s.config.Parent,
		CollectionId: s.config.CollectionId,
	}, ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) Get(r *http.Request, ctx context.Context) (interface{}, error) {
	res, err := s.store.Get(&firestorepb.GetDocumentRequest{}, ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) Create(r *http.Request, ctx context.Context) (interface{}, error) {
	var fields map[string]*firestorepb.Value
	fields["name"] = &firestorepb.Value{ValueType: &firestorepb.Value_StringValue{StringValue: "test"}}

	document := &firestorepb.Document{}

	res, err := s.store.Create(&firestorepb.CreateDocumentRequest{
		Parent:       s.config.Parent,
		CollectionId: s.config.CollectionId,
		Document:     document,
	}, ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) Update(r *http.Request, tx context.Context) (interface{}, error) {
	res, err := s.store.Update(&firestorepb.UpdateDocumentRequest{}, tx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) Delete(r *http.Request, tx context.Context) error {
	err := s.store.Delete(&firestorepb.DeleteDocumentRequest{}, tx)
	if err != nil {
		return err
	}
	return nil
}

func NewService(name string) *Service {
	return &Service{
		Name: name,
	}
}
