package repository

import (
	"access-platform/internal/entity"
	"access-platform/internal/repository/pgdb"
	"access-platform/pkg/postgres"
	"context"
	"github.com/google/uuid"
)

// User defines the methods that any repository implementing user storage must have.
// It allows for saving a new user and retrieving a user by email.
type User interface {
	SaveUser(ctx context.Context, email string, passHash []byte, role string) (uuid.UUID, error)
	GetUser(ctx context.Context, email string) (entity.User, error)
}

type Computer interface {
	AddComputer(ctx context.Context, os string, cpu string, ram int) (uuid.UUID, string, error)
	GetComputer(ctx context.Context, id uuid.UUID) (entity.Computer, error)
	ReserveComputer(ctx context.Context, id uuid.UUID) error
	RelieveComputer(ctx context.Context, id uuid.UUID) error
	GetAllComputers(ctx context.Context) ([]entity.Computer, error)
}

// Repositories is a struct that includes the User repository interface.
// It will hold different repositories for managing various entities in the system.
type Repositories struct {
	User
	Computer
}

// NewRepositories creates a new Repositories instance.
func NewRepositories(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		User:     pgdb.NewUserRepository(pg),
		Computer: pgdb.NewComputerRepository(pg),
	}
}
