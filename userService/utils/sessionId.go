package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"

	"github.com/google/uuid"
)

func GenerateSessionId(email string) string {
	email = strings.ToLower(email)

	// Hash email
	hasher := md5.New()
	hasher.Write([]byte(email))
	hash := hex.EncodeToString(hasher.Sum(nil))

	// Generate UUID4
	uuid4 := uuid.New().String()

	sessionId := hash + "-" + uuid4

	return sessionId
}
