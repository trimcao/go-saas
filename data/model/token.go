package model

import (
	"fmt"
	"strings"

	uuid "github.com/satori/go.uuid"
)

// NewToken returns a token combining an id with a unique identifider
func NewToken(id Key) string {
	newUUID, err := uuid.NewV4()
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%v|%s", id, string(newUUID.String()))
}

// ParseToken returns the id and uuid for a given token
func ParseToken(token string) (string, string) {
	pairs := strings.Split(token, "|")
	if len(pairs) != 2 {
		return "", ""
	}
	return pairs[0], pairs[1]
}
