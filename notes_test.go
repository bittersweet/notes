package main

import (
	"os"
	"testing"
)

func TestGetEditorDefault(t *testing.T) {
	outcome := "vi"
	result := getEditor()
	if result != outcome {
		t.Fatalf("Failed, got %v, expected %v", result, outcome)
	}
}

func TestGetEditorWithEnvVariable(t *testing.T) {
	err := os.Setenv("EDITOR", "nano")
	if err != nil {
		t.Fatalf("Setenv failed")
	}
	outcome := "nano"
	result := getEditor()
	if result != outcome {
		t.Fatalf("Failed, got %v, expected %v", result, outcome)
	}
}

func TestGetHomeDir(t *testing.T) {
	err := os.Setenv("HOME", "/home/testuser")
	if err != nil {
		t.Fatalf("Setenv failed")
	}
	outcome := "/home/testuser/dotfiles/notes/"
	result := getHomeDir()
	if result != outcome {
		t.Fatalf("Failed, got %v, expected %v", result, outcome)
	}
}
