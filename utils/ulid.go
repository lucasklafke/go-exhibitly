package utils

import (
	"crypto/rand" // Use crypto/rand for better entropy
	"time"

	"github.com/oklog/ulid"
)

func GenerateULID() (string, error) { // Return string and error
	ms := ulid.Timestamp(time.Now())
	entropy := rand.Reader // Use crypto/rand.Reader directly

	id, err := ulid.New(ms, entropy) // Capture error
	if err != nil {
		return "", err // Return empty string and the error
	}
	return id.String(), nil // Return the ULID and nil error
}
