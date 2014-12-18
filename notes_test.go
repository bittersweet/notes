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

func TestGetHomeDir(t *testing.T) {
	err := os.Setenv("HOME", "/home/testuser")
	if err != nil {
		t.Fatalf("Setenv failed")
	}
	expected := "/home/testuser/dotfiles/notes/"
	outcome := getHomeDir()
	assert.Equal(t, expected, outcome)
}

func TestGetHomeDirWithEnvVariable(t *testing.T) {
	err := os.Setenv("NOTESDIR", "/some/other/path/")
	if err != nil {
		t.Fatalf("Setenv failed")
	}
	expected := "/some/other/path/"
	outcome := getHomeDir()
	assert.Equal(t, expected, outcome)
}
