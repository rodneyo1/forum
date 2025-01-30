package handlers

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"forum/models"
)

func TestBadRequestHandler_StatusCode400(t *testing.T) {
	w := httptest.NewRecorder()
	BadRequestHandler(w)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestBadRequestHandler_ParseTemplate(t *testing.T) {
	w := httptest.NewRecorder()
	BadRequestHandler(w)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}

	expectedContentType := "text/html; charset=utf-8"
	if contentType := w.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Expected Content-Type %s, got %s", expectedContentType, contentType)
	}

	body := w.Body.String()
	if !strings.Contains(body, "Bad Request!") {
		t.Errorf("Expected response body to contain 'Bad Request!', but it doesn't")
	}
}

func TestBadRequestHandler_SetCodeTo400(t *testing.T) {
	w := httptest.NewRecorder()
	BadRequestHandler(w)

	if hitch.Code != http.StatusBadRequest {
		t.Errorf("Expected hitch.Code to be %d, but got %d", http.StatusBadRequest, hitch.Code)
	}
}

func TestBadRequestHandler_SetIssueToBadRequest(t *testing.T) {
	w := httptest.NewRecorder()
	BadRequestHandler(w)

	if hitch.Issue != "Bad Request!" {
		t.Errorf("Expected hitch.Issue to be 'Bad Request!', but got '%s'", hitch.Issue)
	}
}

func TestBadRequestHandler_ParseTemplateWithHitchData(t *testing.T) {
	w := httptest.NewRecorder()
	BadRequestHandler(w)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}

	expectedHitch := models.WebError{
		Code:  http.StatusBadRequest,
		Issue: "Bad Request!",
	}

	body := w.Body.String()
	if !strings.Contains(body, strconv.Itoa(expectedHitch.Code)) || !strings.Contains(body, expectedHitch.Issue) {
		t.Errorf("Expected response body to contain %d and %s", expectedHitch.Code, expectedHitch.Issue)
	}
}

func TestNotFoundHandler_StatusCode404(t *testing.T) {
	w := httptest.NewRecorder()
	NotFoundHandler(w)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, w.Code)
	}
}

func TestFoundHandler_ParseTemplate(t *testing.T) {
	w := httptest.NewRecorder()
	NotFoundHandler(w)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, w.Code)
	}

	expectedContentType := "text/html; charset=utf-8"
	if contentType := w.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Expected Content-Type %s, got %s", expectedContentType, contentType)
	}

	body := w.Body.String()
	if !strings.Contains(body, "Not Found!") {
		t.Errorf("Expected response body to contain 'Not Found!', but it doesn't")
	}
}

func TestFoundHandler_SetCodeTo404(t *testing.T) {
	w := httptest.NewRecorder()
	NotFoundHandler(w)

	if hitch.Code != http.StatusNotFound {
		t.Errorf("Expected hitch.Code to be %d, but got %d", http.StatusNotFound, hitch.Code)
	}
}

func TestFoundHandler_SetIssueToNotFound(t *testing.T) {
	w := httptest.NewRecorder()
	NotFoundHandler(w)

	if hitch.Issue != "Not Found!" {
		t.Errorf("Expected hitch.Issue to be 'Not Found!', but got '%s'", hitch.Issue)
	}
}

func TestFoundHandler_ParseTemplateWithHitchData(t *testing.T) {
	w := httptest.NewRecorder()
	NotFoundHandler(w)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, w.Code)
	}

	expectedHitch := models.WebError{
		Code:  http.StatusNotFound,
		Issue: "Not Found!",
	}

	body := w.Body.String()
	if !strings.Contains(body, strconv.Itoa(expectedHitch.Code)) || !strings.Contains(body, expectedHitch.Issue) {
		t.Errorf("Expected response body to contain %d and %s", expectedHitch.Code, expectedHitch.Issue)
	}
}
