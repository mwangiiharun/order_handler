package main

import (
	"context"
	"flag"

	"github.com/order_handler/pkg/server"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	ctx := context.Background()

	configFlag := flag.String("config", ".config.yaml", "Path to the config file")
	flag.Parse()

	configService := server.NewConfigProvider(*configFlag)
	config, err := configService.Start(ctx)
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	httpServer, err := server.NewHttpServer(ctx, &config.Server)
	if err != nil {
		log.Fatalf("Failed to create http server: %v", err)
	}

	if err := httpServer.Start(ctx); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
