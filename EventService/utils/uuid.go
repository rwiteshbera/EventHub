package utils

import (
	"github.com/google/uuid"
)

// GenerateUUID : Generate Random UUID
func GenerateUUID() (string, error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return newUUID.String(), nil
}
