package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetApiKey(header http.Header) (string, error) {
	authHeader := header.Get("Authorization")
	apiKey := strings.TrimPrefix(authHeader, "ApiKey ")

	if apiKey == "" {
		return "", fmt.Errorf("API key not present in header")
	}

	return apiKey, nil
}
