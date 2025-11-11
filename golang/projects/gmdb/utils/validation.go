package utils

import (
	"errors"
	"regexp"

	"github.com/google/uuid"
)

// ValidateUUID validates if a string is a valid UUID
func ValidateUUID(uuidStr string) error {
	if _, err := uuid.Parse(uuidStr); err != nil {
		return errors.New("invalid UUID format")
	}
	return nil
}

// ValidateEmail validates if a string is a valid email format
func ValidateEmail(email string) error {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

// ValidateNotEmpty validates if a string is not empty
func ValidateNotEmpty(value string, fieldName string) error {
	if value == "" {
		return errors.New(fieldName + " is required")
	}
	return nil
}
