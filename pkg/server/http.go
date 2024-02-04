package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

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
	log    *logrus.Entry
	Config *HttpConfig
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
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	hs.log.Info("Starting server on port", hs.Config.Port)
	if err := http.ListenAndServe(":"+hs.Config.Port, nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
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
