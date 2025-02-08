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
