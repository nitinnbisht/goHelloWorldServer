package main

import (
	"testing"
	"unicode/utf8"
)

func TestGreetingNotEmpty(t *testing.T) {
	greeting := CreateGreeting("Sam")
	if len(greeting) == 0 {
		t.Error("greeting should not be empty")
	}
}

func TestGreetingValidUTF8(t *testing.T) {
	greeting := CreateGreeting("José")
	if !utf8.ValidString(greeting) {
		t.Error("greeting should contain valid UTF-8")
	}
}

func TestGreetingLength(t *testing.T) {
	greeting := CreateGreeting("X")
	expected := "Hello, X\n"
	if len(greeting) != len(expected) {
		t.Errorf("greeting length mismatch: got %d, want %d", len(greeting), len(expected))
	}
}

func TestMultipleGreetings(t *testing.T) {
	names := []string{"Alice", "Bob", "Charlie"}
	for _, name := range names {
		greeting := CreateGreeting(name)
		if greeting == "" {
			t.Errorf("greeting for %s should not be empty", name)
		}
	}
}
