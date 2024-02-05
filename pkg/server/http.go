package server

import (
	"context"
	"net/http"
	"time"

	"github.com/order_handler/pkg/services"
	"github.com/sirupsen/logrus"
)

type LoggerConfig struct {
	Level logrus.Level `yaml:"level"`
}

type HttpConfig struct {
	Port    string       `yaml:"port"`
	Name    string       `yaml:"name"`
	Version string       `yaml:"version"`
	Logger  LoggerConfig `yaml:"logger"`
}

type HttpServer struct {
	log      *logrus.Entry
	Config   *HttpConfig
	services []services.IService
}

func NewHttpServer(ctx context.Context, config *HttpConfig) (*HttpServer, error) {
	return &HttpServer{
		Config: config,
	}, nil
}

func (hs *HttpServer) Start(ctx context.Context) error {
	err := hs.configureLogging(ctx)
	if err != nil {
		return err
	}

	
	for _, service := range hs.services {
		if service.GetStatus() == services.Running {
			hs.log.Info("Registering handlers for service ", service.GetName())
			handlers := service.GetHandlers()
			for path, handler := range handlers {
				hs.log.Info("Registering handler for ", path)
				http.HandleFunc(path, handler)
			}
		}
	}

	hs.log.Info("Starting server on port ", hs.Config.Port)
	if err := http.ListenAndServe(":"+hs.Config.Port, nil); err != nil {
		return err
	}
	return nil
}

func (hs *HttpServer) configureLogging(ctx context.Context) error {
	log := logrus.New()
	globalFields := logrus.Fields{
		"app":     hs.Config.Name,
		"version": hs.Config.Version,
		"time":    time.Now().Format(time.RFC3339),
	}
	entry := log.WithFields(globalFields)
	entry.Logger.SetLevel(hs.Config.Logger.Level)
	hs.log = entry

	return nil
}
