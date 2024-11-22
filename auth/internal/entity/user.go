package entity

import "github.com/google/uuid"

// User represents the structure for a general user in the system.
// It includes fields for identifying, authenticating, and defining the user's role.
type User struct {
	ID           uuid.UUID
	Email        string
	HashPassword []byte
	Role         string
}
