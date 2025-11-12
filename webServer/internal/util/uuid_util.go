package util

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	uuidStr := uuid.New().String()
	uuidStr = strings.ReplaceAll(uuidStr, "-", "")
	return uuidStr[:20]
}
