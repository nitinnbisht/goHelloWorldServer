package main

import (
	"strings"
	"testing"
)

func TestEdgeCaseLongName(t *testing.T) {
	longName := strings.Repeat("A", 1000)
	greeting := CreateGreeting(longName)
	if !strings.Contains(greeting, longName) {
		t.Error("greeting should contain long name")
	}
}

func TestEdgeCaseSpecialCharacters(t *testing.T) {
	names := []string{"O'Brien", "Jean-Paul", "María", "李明", "😀"}
	for _, name := range names {
		greeting := CreateGreeting(name)
		if !strings.Contains(greeting, name) {
			t.Errorf("greeting should contain %q", name)
		}
	}
}

func TestEdgeCaseWhitespace(t *testing.T) {
	greeting := CreateGreeting("\t\n  ")
	if len(greeting) == 0 {
		t.Error("greeting with whitespace should not be empty")
	}
}

func TestEdgeCaseNumbers(t *testing.T) {
	greeting := CreateGreeting("123")
	if !strings.Contains(greeting, "123") {
		t.Error("greeting should contain numbers")
	}
}
