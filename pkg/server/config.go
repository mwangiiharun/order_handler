package server

import (
	"context"
	"os"

	"github.com/order_handler/pkg/services"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server   HttpConfig               `yaml:"server"`
	Storage  services.FirestoreServiceConfig    `yaml:"storage"`
	Services []services.ServiceConfig `yaml:"services"`
}

type ConfigService struct {
	filePath string
}

func NewConfigProvider(filePath string) *ConfigService {
	return &ConfigService{filePath: filePath}
}

func (cs *ConfigService) Start(ctx context.Context) (*Config, error) {
	data, err := os.ReadFile(cs.filePath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
