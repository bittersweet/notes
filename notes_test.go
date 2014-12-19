package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetEditorDefault(t *testing.T) {
	expected := "vi"
	outcome := getEditor()
	assert.Equal(t, expected, outcome)
}

func TestGetEditorWithEnvVariable(t *testing.T) {
	err := os.Setenv("EDITOR", "nano")
	if err != nil {
		t.Fatalf("Setenv failed")
	}
	expected := "nano"
	outcome := getEditor()
	assert.Equal(t, expected, outcome)
}

func TestGetNotesDir(t *testing.T) {
	err := os.Setenv("HOME", "/home/testuser")
	if err != nil {
		t.Fatalf("Setenv failed")
	}
	expected := "/home/testuser/dotfiles/notes/"
	outcome := getNotesDir()
	assert.Equal(t, expected, outcome)
}

func TestGetNotesDirWithCustomEnvVariable(t *testing.T) {
	err := os.Setenv("NOTESDIR", "/some/other/path/")
	if err != nil {
		t.Fatalf("Setenv failed")
	}
	expected := "/some/other/path/"
	outcome := getNotesDir()
	assert.Equal(t, expected, outcome)
}

func TestGetNotesDirUsesCustomEnvVariableOverHomeEnvVariable(t *testing.T) {
	err := os.Setenv("HOME", "/home/testuser")
	if err != nil {
		t.Fatalf("Setenv failed")
	}
	err = os.Setenv("NOTESDIR", "/some/other/path/")
	if err != nil {
		t.Fatalf("Setenv failed")
	}
	expected := "/some/other/path/"
	outcome := getNotesDir()
	assert.Equal(t, expected, outcome)
}
