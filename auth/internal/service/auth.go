package service

import (
	"access-platform/internal/entity"
	"access-platform/internal/lib/jwt"
	"access-platform/internal/repository"
	"access-platform/internal/repository/pgdb"
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

// Auth defines the methods that an authentication service should implement.
type Auth interface {
	Register(ctx context.Context, email string, password string, role string) (entity.User, error)
	Login() error
}

// AuthDependencies holds dependencies for the AuthService, including logger, repository, and secret key.
type AuthDependencies struct {
	Log       *zap.Logger
	Repos     *repository.Repositories
	SecretKey string
}

// AuthService is a service for handling user authentication, including registration and login.
type AuthService struct {
	log            *zap.Logger
	userRepository repository.User
	secretKey      string
}

// New creates a new instance of AuthService with the provided dependencies.

func New(deps AuthDependencies) *AuthService {
	return &AuthService{
		log:            deps.Log,
		userRepository: deps.Repos.User,
		secretKey:      deps.SecretKey,
	}
}

// Register handles the user registration process. It hashes the password, saves the user, and returns the registered user.
func (s *AuthService) Register(ctx context.Context, email string, password string, role string) (entity.RegisteredUser, error) {
	const op = "service.Auth.Register"

	s.log.Info("attempting to register user", zap.String("op", op))

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		s.log.Error("failed to generate password hash", zap.Error(err), zap.String("op", op))

		return entity.RegisteredUser{}, fmt.Errorf("%s: %w", op, err)
	}

	id, err := s.userRepository.SaveUser(ctx, email, passwordHash, role)
	if err != nil {
		if errors.Is(err, pgdb.ErrUserAlreadyExists) {
			s.log.Warn("user already exists", zap.Error(err), zap.String("op", op))

			return entity.RegisteredUser{}, fmt.Errorf("%s: %w", op, ErrUserAlreadyExists)
		}
	}

	registeredUser := entity.RegisteredUser{
		ID:    id,
		Email: email,
		Role:  role,
	}

	return registeredUser, nil
}

// Login handles the login process by verifying the user's credentials and generating a JWT token if successful.
func (s *AuthService) Login(ctx context.Context, email string, password string) (entity.LoggedUser, error) {
	const op = "service.Auth.Login"

	s.log.Info("attempting to login user", zap.String("op", op))

	user, err := s.userRepository.GetUser(ctx, email)
	if err != nil {
		if errors.Is(err, pgdb.ErrUserNotFound) {
			s.log.Warn("user not found", zap.Error(err), zap.String("op", op))

			return entity.LoggedUser{}, fmt.Errorf("%s: %w", op, ErrUserNotFound)
		}

		return entity.LoggedUser{}, fmt.Errorf("%s: %w", op, err)
	}

	if err := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password)); err != nil {
		s.log.Info("invalid credentials", zap.Error(err), zap.String("op", op))

		return entity.LoggedUser{}, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	s.log.Info("user logged successfully", zap.String("op", op))

	token, err := jwt.NewToken(user.ID, user.Email, user.PasswordHash, user.Role, s.secretKey)
	if err != nil {
		s.log.Error("failed to generate token", zap.Error(err), zap.String("op", op))

		return entity.LoggedUser{}, fmt.Errorf("%s: %w", op, err)
	}

	loggedUser := entity.LoggedUser{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
		Token: token,
	}

	return loggedUser, nil
}
