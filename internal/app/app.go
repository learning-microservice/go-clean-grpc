package app

import (
	"fmt"
	"os"

	"go-clean-grpc/config"
	"go-clean-grpc/internal/delivery/grpc"
	"go-clean-grpc/internal/registry"
	"go-clean-grpc/pkg/grpcserver"
	"go-clean-grpc/pkg/log"
)

// Run -.
func Run(cfg *config.Config) error {
	// setup Logger
	logger := log.New(os.Stdout,
		log.Level(cfg.Log.Level),
		log.GlobalAttrs(
			"app", cfg.APP.Name,
			"version", cfg.APP.Version,
		),
	)

	// setup registry
	reg, err := registry.New(cfg)
	if err != nil {
		return fmt.Errorf("failed to setup registry: %w", err)
	}

	// setup rest-api server
	server := grpcserver.New(
		grpcserver.Address(cfg.HTTP.Host, cfg.HTTP.Port),
		grpcserver.Logger(logger),
	)

	// setup router
	grpc.SetupRouter(server, reg)

	return server.ListenAndServe()
}
