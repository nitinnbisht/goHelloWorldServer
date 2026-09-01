package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIntegrationMultipleRequests(t *testing.T) {
	requests := []string{"/?name=User1", "/?name=User2", "/"}

	for _, reqPath := range requests {
		req := httptest.NewRequest("GET", reqPath, nil)
		w := httptest.NewRecorder()

		handler(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("request %s returned %d, expected 200", reqPath, w.Code)
		}
	}
}

func TestIntegrationQueryParsing(t *testing.T) {
	tests := []struct {
		path     string
		hasName  bool
		contains string
	}{
		{"/?name=TestA", true, "TestA"},
		{"/?name=", false, "Guest"},
		{"/", false, "Guest"},
	}

	for _, tt := range tests {
		req := httptest.NewRequest("GET", tt.path, nil)
		w := httptest.NewRecorder()

		handler(w, req)

		body := w.Body.String()
		if !contains(body, tt.contains) {
			t.Errorf("path %s: expected %q in %q", tt.path, tt.contains, body)
		}
	}
}

func TestIntegrationHTTPMethods(t *testing.T) {
	req := httptest.NewRequest("GET", "/?name=Test", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("GET request failed with status %d", w.Code)
	}
}

func contains(s, substr string) bool {
	for i := 0; i < len(s)-len(substr)+1; i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
