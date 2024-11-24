package service

import (
	"access-platform/internal/entity"
	"access-platform/internal/repository"
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ComputeAccess interface {
	AddComputer(ctx context.Context, os string, cpu string, ram int) (entity.Computer, error)
	GetComputer(ctx context.Context, id uuid.UUID) (entity.Computer, error)
	ReserveComputer(ctx context.Context, id uuid.UUID) error
	RelieveComputer(ctx context.Context, id uuid.UUID) error
	GetAllComputers(ctx context.Context) ([]entity.Computer, error)
}

type ComputeAccessDependencies struct {
	Log   *zap.Logger
	Repos *repository.Repositories
}

type ComputeAccessService struct {
	log                *zap.Logger
	computerRepository repository.Computer
}

func NewComputeAccessService(deps ComputeAccessDependencies) *ComputeAccessService {
	return &ComputeAccessService{
		log:                deps.Log,
		computerRepository: deps.Repos.Computer,
	}
}

func (s *ComputeAccessService) AddComputer(ctx context.Context, os string, cpu string, ram int) (entity.Computer, error) {
	const op = "service.Access.AddComputer"

	s.log.Info("attempting to add computer", zap.String("op", op))

	id, ssh, err := s.computerRepository.AddComputer(ctx, os, cpu, ram)
	if err != nil {
		s.log.Error("failed to add computer", zap.Error(err), zap.String("op", op))

		return entity.Computer{}, fmt.Errorf("%s: %w", op, err)
	}

	s.log.Info("computer added successfully", zap.String("op", op))

	computer := entity.Computer{
		ID:     id,
		OS:     os,
		CPU:    cpu,
		RAM:    ram,
		Status: true,
		SSH:    ssh,
	}

	return computer, nil
}

func (s *ComputeAccessService) GetComputer(ctx context.Context, id uuid.UUID) (entity.Computer, error) {
	const op = "service.Access.GetComputer"

	s.log.Info("attempting to get computer", zap.String("op", op))

	computer, err := s.computerRepository.GetComputer(ctx, id)
	if err != nil {
		s.log.Error("failed to get computer", zap.Error(err), zap.String("op", op))

		return entity.Computer{}, fmt.Errorf("%s: %w", op, err)
	}

	return computer, nil
}

func (s *ComputeAccessService) ReserveComputer(ctx context.Context, id uuid.UUID) error {
	const op = "service.Access.ReserveComputer"

	s.log.Info("attempting to reserve computer", zap.String("op", op))

	err := s.computerRepository.ReserveComputer(ctx, id)
	if err != nil {
		s.log.Error("failed to reserve computer", zap.Error(err), zap.String("op", op))

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *ComputeAccessService) RelieveComputer(ctx context.Context, id uuid.UUID) error {
	const op = "service.Access.RelieveComputer"

	s.log.Info("attempting to relieve computer", zap.String("op", op))

	err := s.computerRepository.RelieveComputer(ctx, id)
	if err != nil {
		s.log.Error("failed to relieve computer", zap.Error(err), zap.String("op", op))

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *ComputeAccessService) GetAllComputers(ctx context.Context) ([]entity.Computer, error) {
	const op = "service.Access.GetAllComputers"

	s.log.Info("attempting to get all computers", zap.String("op", op))

	// Вызов метода репозитория для получения всех компьютеров
	computers, err := s.computerRepository.GetAllComputers(ctx)
	if err != nil {
		s.log.Error("failed to get all computers", zap.Error(err), zap.String("op", op))

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	s.log.Info("successfully retrieved all computers", zap.Int("count", len(computers)), zap.String("op", op))

	return computers, nil
}
