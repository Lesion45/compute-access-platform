package pgdb

import (
	"access-platform/internal/entity"
	"access-platform/pkg/postgres"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
)

// UserRepository is a concrete implementation of the User repository interface.
// It provides methods to save and retrieve user data from a PostgreSQL database.
type UserRepository struct {
	*postgres.Postgres
}

// NewUserRepository creates a new instance of UserRepository.
func NewUserRepository(pg *postgres.Postgres) *UserRepository {
	return &UserRepository{pg}
}

// SaveUser saves a new user to the database. It inserts the user's email, hashed password, and role.
// If the user already exists, it returns an error indicating so.
func (r *UserRepository) SaveUser(ctx context.Context, email string, passHash []byte, role string) (uuid.UUID, error) {
	const op = "repository.user.SaveUser"

	var userID uuid.UUID

	query := `INSERT INTO users_schema.user(email, password_hash, role) VALUES(@userEmail, @userPasswordHash, @userRole) RETURNING id`
	args := pgx.NamedArgs{
		"userEmail":        email,
		"userPasswordHash": passHash,
		"userRole":         role,
	}

	err := r.DB.QueryRow(ctx, query, args).Scan(&userID)
	if err != nil {
		pgErr, ok := err.(*pgconn.PgError)
		if ok {
			if pgErr.Code == "23505" {
				return uuid.Nil, fmt.Errorf("%s: %w", op, ErrUserAlreadyExists)
			}
		} else {
			return uuid.Nil, fmt.Errorf("%s: %w", op, err)
		}
	}
	fmt.Println(err)

	return userID, err
}

// GetUser retrieves a user from the database by their email address.
// If the user is found, it returns the user entity; otherwise, it returns an error.
func (r *UserRepository) GetUser(ctx context.Context, email string) (entity.User, error) {
	const op = "repository.user.GetUser"

	var userID uuid.UUID
	var userEmail string
	var userPasswordHash []byte
	var userRole string

	query := `SELECT id, email, password_hash, role FROM users_schema.user WHERE email = @userEmail`
	args := pgx.NamedArgs{
		"userEmail": email,
	}

	err := r.DB.QueryRow(ctx, query, args).Scan(&userID, &userEmail, &userPasswordHash, &userRole)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, fmt.Errorf("%s: %w", op, ErrUserNotFound)
		}

		return entity.User{}, fmt.Errorf("%s: %w", op, err)
	}

	user := entity.User{
		ID:           userID,
		Email:        userEmail,
		PasswordHash: userPasswordHash,
		Role:         userRole,
	}

	return user, nil
}
