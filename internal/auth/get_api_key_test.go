package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey abc123")

	got, err := GetAPIKey(headers)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	want := "abc123"
	if got != want {
		t.Fatalf("expected API key %q, got %q", want, got)
	}
}

func TestGetAPIKeyNoHeader(t *testing.T) {
	headers := http.Header{}

	got, err := GetAPIKey(headers)

	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("expected ErrNoAuthHeaderIncluded, got: %v", err)
	}

	if got != "" {
		t.Fatalf("expected empty API key, got %q", got)
	}
}

func TestGetAPIKeyMalformedHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer abc123")

	got, err := GetAPIKey(headers)

	if err == nil {
		t.Fatal("expected an error for malformed authorization header")
	}

	if got != "" {
		t.Fatalf("expected empty API key, got %q", got)
	}
}
