package models

// UserRegister represents the structure for user registration data.
// It contains the email and password required for creating a new user account.
type UserRegister struct {
	Email    string `json:"user-email"`
	Password string `json:"user-password"`
}

// UserLogin represents the structure for user login data.
// It contains the email and password required to authenticate an existing user.
type UserLogin struct {
	Email    string `json:"user-email"`
	Password string `json:"user-password"`
}
