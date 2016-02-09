package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestFindNotes(t *testing.T) {
	var notes []Note
	notes = append(notes, Note{
		Explanation: []string{"# Explanation first"},
		Command:     []string{"./notes first"},
	})
	notes = append(notes, Note{
		Explanation: []string{"# Explanation second"},
		Command:     []string{"./notes second"},
	})
	result := findNotes(notes, "second")
	assert.Equal(t, 1, len(result))
	assert.Equal(t, result[0].Command, []string{"./notes second"})
}

func TestFindNotesSearchesCommandsAsWell(t *testing.T) {
	var notes []Note
	notes = append(notes, Note{
		Explanation: []string{"# Explanation first"},
		Command:     []string{"./notes first"},
	})
	notes = append(notes, Note{
		Explanation: []string{"# Explanation second"},
		Command:     []string{"./notes match"},
	})
	result := findNotes(notes, "match")
	assert.Equal(t, 1, len(result))
	assert.Equal(t, result[0].Command, []string{"./notes match"})
}

func TestFindNotesCaseInsensitivity(t *testing.T) {
	var notes []Note

	note := Note{
		Explanation: []string{"# Explanation note MATCH"},
		Command:     []string{"Run rm -rf", "echo 'done'"},
	}
	notes = append(notes, note)
	result := findNotes(notes, "match")
	assert.Equal(t, 1, len(result))
}
