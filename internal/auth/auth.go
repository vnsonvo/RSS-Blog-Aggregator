package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GETAPIKey extracts an API Key from the headers of the HTTP request
// Example: Authorization: ApiKey {apikey here}
func GetAPIKey(headers http.Header) (string, error) {
	reqToken := headers.Get("Authorization")
	if reqToken == "" {
		return "", errors.New("no authentication info found")
	}

	splitToken := strings.Split(reqToken, " ")
	if len(splitToken) != 2 || splitToken[0] != "ApiKey" {
		return "", errors.New("bad auth header")
	}
	return splitToken[1], nil
}
