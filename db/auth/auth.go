package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Example:
// Authorization: ApiKey YOUR_API_KEY

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no Authorization header found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed Authorization header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("unsupported authorization type")
	}

	return vals[1], nil
}
