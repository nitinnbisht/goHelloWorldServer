package main

import (
	"strings"
	"testing"
)

func TestStringUtils(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty", "", "Hello, Guest\n"},
		{"with spaces", "  Alice  ", "Hello,   Alice  \n"},
		{"special chars", "Bob@123", "Hello, Bob@123\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CreateGreeting(tt.input)
			if strings.TrimSpace(result) != strings.TrimSpace(tt.expected) {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestGreetingFormat(t *testing.T) {
	greeting := CreateGreeting("Test")
	if !strings.HasPrefix(greeting, "Hello,") {
		t.Errorf("greeting should start with 'Hello,', got: %s", greeting)
	}
	if !strings.HasSuffix(greeting, "\n") {
		t.Errorf("greeting should end with newline, got: %s", greeting)
	}
}
