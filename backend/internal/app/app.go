package app

import (
	"access-platform/config"
	auth "access-platform/internal/controller/http"
	"access-platform/internal/repository"
	"access-platform/internal/service"
	"access-platform/pkg/logger"
	"access-platform/pkg/postgres"
	"context"
	"go.uber.org/zap"
	"net/http"
	"os"
)

func Run() {
	cfg := config.MustLoad()

	log := logger.NewZap(cfg.Env)

	log.Info("Initializing application...")
	pg, err := postgres.NewPG(context.Background(), cfg.DB)
	if err != nil {
		log.Error("database initialization error", zap.Error(err))
		os.Exit(1)
	}

	err = pg.PostgresHealthCheck(context.Background())
	if err != nil {
		log.Error("Postgres doesn't response", zap.Error(err))
		os.Exit(1)
	}
	defer pg.DB.Close()
	log.Info("Initializing postgres: successful!")

	log.Info("Initializing repositories...")
	repositories := repository.NewRepositories(pg)
	log.Info("Initializing repositories: successful!")

	log.Info("Initializing services...")
	deps := service.ServiceDependencies{
		Log:       log,
		Repos:     repositories,
		SecretKey: cfg.SecretKey,
	}
	services := service.New(deps)
	log.Info("Initializing services: successful!")

	router := auth.InitRouter(log, services)

	server := &http.Server{
		Addr:         cfg.Server.Addr,
		Handler:      router,
		ReadTimeout:  cfg.Server.TimeOut,
		WriteTimeout: cfg.Server.TimeOut,
		IdleTimeout:  cfg.Server.IdleTimeOut,
	}
	log.Info("Initializing server: OK!")

	if err := server.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}
	log.Error("server stopped")
}
