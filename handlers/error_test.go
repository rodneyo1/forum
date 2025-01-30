package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBadRequestHandler(t *testing.T) {
	w := httptest.NewRecorder()
	BadRequestHandler(w)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}
}
