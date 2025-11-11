package utils

import (
	"github.com/google/uuid"
)

// NewUUIDv7 generates a new UUIDv7 (time-ordered UUID)
// UUIDv7 provides better database performance due to sequential ordering
func NewUUIDv7() uuid.UUID {
	v7, err := uuid.NewV7()
	if err != nil {
		// Fallback to UUIDv4 if v7 generation fails
		return uuid.New()
	}
	return v7
}

// NewUUID is an alias for backward compatibility
// Use NewUUIDv7() for new code
func NewUUID() uuid.UUID {
	return NewUUIDv7()
}
