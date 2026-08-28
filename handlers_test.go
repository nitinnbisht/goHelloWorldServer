package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlerWithName(t *testing.T) {
	req := httptest.NewRequest("GET", "/?name=Alice", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

func TestHandlerWithoutName(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

func TestHandlerResponseBody(t *testing.T) {
	req := httptest.NewRequest("GET", "/?name=TestUser", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	body := w.Body.String()
	if !strings.Contains(body, "TestUser") {
		t.Errorf("expected body to contain 'TestUser', got: %s", body)
	}
}
