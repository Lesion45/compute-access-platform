package entity

import "github.com/google/uuid"

type Computer struct {
	ID     uuid.UUID
	OS     string
	CPU    string
	RAM    int
	Status bool
	SSH    string
}
