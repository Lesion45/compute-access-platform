package auth

import (
	"access-platform/internal/controller/http/response"
	"access-platform/internal/service"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
)

// RegisterRequest is the structure that holds the incoming registration request
// containing the user's email, password, and role.
type RegisterRequest struct {
	Email    string `json:"user-email"`
	Password string `json:"user-password"`
	Role     string `json:"user-role"`
}

// RegisterResponse is the structure returned in the response after a successful user registration.
// It contains the user ID, email, and role.
type RegisterResponse struct {
	ID    uuid.UUID `json:"user-id"`
	Email string    `json:"user-email"`
	Role  string    `json:"user-role"`
}

// Register is the handler function for registering a user in the system.
// It accepts the user details in the request body, interacts with the AuthService,
// and returns the user information on successful registration.
func Register(log *zap.Logger, auth *service.AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const op = "handlers.v1.Register"

		var req RegisterRequest

		err := ctx.BindJSON(&req)
		if err != nil {
			log.Error("failed to decode request body", zap.Error(err), zap.String("op", op))
			ctx.IndentedJSON(http.StatusInternalServerError, response.Error("failed to decode request"))

			return
		}

		log.Info("request body decoded", zap.Any("request", req), zap.String("op", op))

		user, err := auth.Register(ctx, req.Email, req.Password, req.Role)
		if err != nil {
			log.Info("failed to register user", zap.String("op", op))

			if errors.Is(err, service.ErrUserAlreadyExists) {
				ctx.IndentedJSON(http.StatusConflict, response.Error("user already exists"))

				return
			}
			ctx.IndentedJSON(http.StatusInternalServerError, response.Error("internal server error"))

			return
		}

		log.Info("user registered", zap.String("op", op))

		ctx.IndentedJSON(http.StatusOK, RegisterResponse{
			ID:    user.ID,
			Email: user.Email,
			Role:  user.Role,
		})

		return
	}
}

type LoginRequest struct {
	Email    string `json:"user-email"`
	Password string `json:"user-password"`
}

type LoginResponse struct {
	ID    uuid.UUID `json:"user-id"`
	Email string    `json:"user-email"`
	Role  string    `json:"user-role"`
	Token string    `json:"token"`
}

func Login(log *zap.Logger, auth *service.AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const op = "handlers.v1.Login"

		var req LoginRequest

		err := ctx.BindJSON(&req)
		if err != nil {
			log.Error("failed to decode request body", zap.Error(err), zap.String("op", op))
			ctx.IndentedJSON(http.StatusInternalServerError, response.Error("failed to decode request"))

			return
		}

		log.Info("request body decoded", zap.Any("request", req), zap.String("op", op))

		user, err := auth.Login(ctx, req.Email, req.Password)
		if err != nil {
			log.Info("failed to login", zap.String("op", op))

			if errors.Is(err, service.ErrInvalidCredentials) {
				ctx.IndentedJSON(http.StatusUnauthorized, response.Error("invalid credentials"))

				return
			} else if errors.Is(err, service.ErrUserNotFound) {
				ctx.IndentedJSON(http.StatusNotFound, response.Error("user doesn't exist"))

				return
			}
			ctx.IndentedJSON(http.StatusInternalServerError, response.Error("internal server error"))

			return
		}

		log.Info("user logged in", zap.String("op", op))

		ctx.IndentedJSON(http.StatusOK, LoginResponse{
			ID:    user.ID,
			Email: user.Email,
			Role:  user.Role,
			Token: user.Token,
		})

		return
	}
}
