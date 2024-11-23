package service

import (
	"access-platform/internal/repository"
	"go.uber.org/zap"
)

type ServiceDependencies struct {
	Log       *zap.Logger
	Repos     *repository.Repositories
	SecretKey string
}

type Service struct {
	log            *zap.Logger
	AuthService    *AuthService
	ComputeService *ComputeAccessService
	SecretKey      string
}

func New(deps ServiceDependencies) *Service {
	return &Service{
		log: deps.Log,
		AuthService: NewAuthService(AuthDependencies{
			Log:       deps.Log,
			Repos:     deps.Repos,
			SecretKey: deps.SecretKey,
		}),
		ComputeService: NewComputeAccessService(ComputeAccessDependencies{
			Log:   deps.Log,
			Repos: deps.Repos,
		}),
	}
}
