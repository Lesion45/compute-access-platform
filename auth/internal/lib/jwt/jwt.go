package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

// Claims defines the custom claims for the JWT token.
type Claims struct {
	UID      uuid.UUID `json:"uid"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Role     string    `json:"role"`
	jwt.RegisteredClaims
}

// NewToken generates a new JWT token for a given user.
//
// Parameters:
//   - uid: Unique user ID (UUID).
//   - email: User's email address.
//   - hashPassword: Hashed user password (byte array).
//   - role: User's role ("admin" or "user").
//   - secretKey: Secret key used for signing the JWT token.
//
// Returns:
//   - A signed JWT token string.
//   - An error if the token creation fails.
func NewToken(uid uuid.UUID, email string, hashPassword []byte, role string, secretKey string) (string, error) {
	const op = "jwt.NewToken"

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UID:      uid,
		Email:    email,
		Password: string(hashPassword),
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return tokenString, nil
}
