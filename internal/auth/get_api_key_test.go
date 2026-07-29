package auth

import (
	"net/http"
	"testing"
)

// var ErrNoAuthHeaderIncluded = errors.New("no authorization header included")

// GetAPIKey -
func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}

	_, err := GetAPIKey(headers)
	if err == ErrNoAuthHeaderIncluded {
		t.Fatalf("expected error %v, got %v", ErrNoAuthHeaderIncluded, err)
	}
}
