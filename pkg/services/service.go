package services

import (
	"context"
	"fmt"
	"net/http"
)

type ServiceType int

const (
	Storage ServiceType = iota + 1
	Orders
	Customers
	Products
	Payments
)

func (s ServiceType) String() string {
	return [...]string{"Storage", "Orders", "Customers", "Products", "Payments"}[s-1]
}

func (s ServiceType) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}

func (s ServiceType) UnmarshalText(text []byte) error {
	switch string(text) {
	case "Storage":
		s = Storage
	case "Orders":
		s = Orders
	case "Customers":
		s = Customers
	case "Products":
		s = Products
	case "Payments":
		s = Payments
	default:
		return fmt.Errorf("invalid ServiceType: %s", text)
	}
	return nil
}

type serviceStatus int

const (
	Stopped serviceStatus = iota + 1
	Running
)

type IService interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error

	GetHandlers() map[string]func(http.ResponseWriter, *http.Request)
	GetName() string
	GetStatus() serviceStatus
}

type KV struct {
	Key   string
	Value string
}

type ServiceConfig struct {
	Name string      `yaml:"name"`
	Type ServiceType `yaml:"type"`
	meta KV
}

type ServiceConfigs *map[string]ServiceConfig

type Service struct {
	config   *ServiceConfig
	status   serviceStatus
	handlers map[string]func(http.ResponseWriter, *http.Request)
}

func (s *Service) GetName() string {
	return s.config.Name
}

func (s *Service) GetStatus() serviceStatus {
	return s.status
}

func (s *Service) Start(ctx context.Context) error {
	if s.config == nil {
		return fmt.Errorf("service config not found")
	}
	s.status = Running
	return nil
}

func (s *Service) Stop(ctx context.Context) error {
	if s.status == Stopped {
		return fmt.Errorf("service already stopped")
	}

	s.status = Stopped
	return nil
}

func (s *Service) GetHandlers() map[string]func(http.ResponseWriter, *http.Request) {
	return s.handlers
}

func RegisterServices(ctx context.Context, configs ServiceConfigs) (map[string]IService, error) {
	services := make(map[string]IService)
	for name, config := range *configs {
		switch config.Type {
		case Storage:
			services[name] = NewFirestore(ctx, config)
		case Orders:
			services[name] = NewOrdersService(ctx, config)
		default:
			return nil, fmt.Errorf("invalid service type: %s", config.Type)
		}
	}
	return services, nil
}
