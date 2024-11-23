package entity

import "github.com/google/uuid"

// User represents the structure for a general user in the system.
// It includes fields for identifying, authenticating, and defining the user's role.
type User struct {
	ID           uuid.UUID
	Email        string
	PasswordHash []byte
	Role         string
}

// RegisteredUser represents a user who has been registered in the system
// but has not yet authenticated. It contains the user's unique ID, email,
// and role. This structure is typically used during the user registration
type RegisteredUser struct {
	ID    uuid.UUID
	Email string
	Role  string
}

// LoggedUser represents an authenticated user. It contains all the fields
// from RegisteredUser, along with a JWT token that is issued upon successful
// authentication.
type LoggedUser struct {
	ID    uuid.UUID
	Email string
	Role  string
	Token string
}
