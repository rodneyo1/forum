package database

import (
	"net/http"
	"testing"
)

func TestIsLoggedIn(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	result := IsLoggedIn(req)

	if result {
		t.Errorf("IsLoggedIn() = %v, want %v", result, false)
	}
}

func TestIsLoggedIn_EmptySessionCookie(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set an empty session cookie
	cookie := &http.Cookie{
		Name:  "session_id",
		Value: "",
	}
	req.AddCookie(cookie)

	result := IsLoggedIn(req)
	if result {
		t.Errorf("IsLoggedIn() returned true for empty session cookie, expected false")
	}
}
