package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	Email    string
	Password string
	Role     string
	jwt.RegisteredClaims
}

// NewToken creates new JWT token for given user.
func NewToken(email string, hashPassword []byte, role string, duration time.Duration, secretKey string) (string, error) {
	expirationTime := time.Now().Add(duration)

	claims := &Claims{
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
		return "", err
	}

	return tokenString, nil
}
