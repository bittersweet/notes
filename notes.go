package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/codegangsta/cli"
)

// Note holds parsed explanation and commands
type Note struct {
	Explanation []string
	Command     []string
}

// Print outputs an indivial note comment and command to STDOUT
func (n *Note) Print() {
	for _, line := range n.Explanation {
		colorizeComment(line)
	}

	for _, line := range n.Command {
		fmt.Println(line)
	}
	fmt.Println()
}

func findNotes(notes []Note, query string) []Note {
	var results []Note

Loop:
	for _, note := range notes {
		queryRegexp := fmt.Sprintf("(?i)%s", query)
		searchRegexp, _ := regexp.Compile(queryRegexp)

		for _, line := range note.Explanation {
			if searchRegexp.MatchString(line) == true {
				results = append(results, note)
				// We already have a match, continue with the next note
				continue Loop
			}
		}

		for _, line := range note.Command {
			if searchRegexp.MatchString(line) == true {
				results = append(results, note)
			}
		}
	}

	return results
}

func showAllNotes() {
	path := fmt.Sprintf("%v/*.txt", getNotesDir())
	matches, err := filepath.Glob(path)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	for i := 0; i < len(matches); i++ {
		match := matches[i]
		filename := filepath.Base(match)
		baseSize := len(filename) - 4 // Strip .txt from note filenames
		fmt.Println(filename[:baseSize])
	}
}

func showNote(params ...string) {
	note := params[0]
	var query string
	// params will be > 1 if a search term has been entered:
	// $ notes <note> <query>
	if len(params) > 1 {
		query = params[1]
	}

	path := getNote(note)
	file, err := os.Open(path)
	if err != nil {
		// File does not exist, create it
		editOrCreateNote(note)
		return
	}
	defer file.Close()

	var notes []Note
	var n Note

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		commentRegexp, _ := regexp.Compile("^#")
		if commentRegexp.MatchString(line) == true {
			n.Explanation = append(n.Explanation, line)
		} else if line == "" {
			// newline means the start of a new note so we add the last found note
			// and set n to a new note to restart the process
			notes = append(notes, n)
			n = Note{} // Reset note
		} else {
			n.Command = append(n.Command, line)
		}
	}
	if len(n.Explanation) > 0 {
		// Do not append an empty note if a newline was the last line of a file
		notes = append(notes, n)
	}

	if len(params) > 1 {
		notes = findNotes(notes, query)
	}

	for i := 0; i < len(notes); i++ {
		notes[i].Print()
	}
}

func colorizeComment(line string) {
	highlight := "\033[33m"
	reset := "\033[0m"

	fmt.Printf("%v%v%v\n", highlight, line, reset)
}

func editOrCreateNote(note string) {
	path := getNote(note)

	command := exec.Command(getEditor(), path)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Run()
}

func getEditor() string {
	editor := os.Getenv("EDITOR")
	if len(editor) == 0 {
		editor = "vi"
	}
	return editor
}

// getNotesDir returns the path to the configured notes
// dir, without a trailing slash
func getNotesDir() string {
	notesdir := os.Getenv("NOTESDIR")
	var dir string
	if notesdir != "" {
		dir = notesdir
	} else {
		// Use default location, ~/dotfiles/notes/
		homedir := os.Getenv("HOME")
		dir = fmt.Sprintf("%v/dotfiles/notes/", homedir)
	}

	return filepath.Clean(dir)
}

// getNote returns the path to the note given
func getNote(path string) string {
	return fmt.Sprintf("%s/%s.txt", getNotesDir(), path)
}

func main() {
	app := cli.NewApp()
	app.Name = "notes"
	app.Version = "0.6.1"
	app.Usage = "Store your thoughts on all sorts of subjects"
	app.Action = func(c *cli.Context) {
		note := c.Args().First()
		if len(c.Args()) == 0 {
			// $ notes
			showAllNotes()
		} else if len(c.Args()) == 1 {
			// $ notes <note>
			showNote(note)
		} else {
			// $ notes <note> <query>
			showNote(note, c.Args()[1])
		}
	}
	app.Commands = []cli.Command{
		{
			Name:      "list",
			ShortName: "l",
			Usage:     "List all notes",
			Action: func(c *cli.Context) {
				// $ notes list
				showAllNotes()
			},
		},
		{
			Name:      "new",
			ShortName: "n",
			Usage:     "Create new note",
			Action: func(c *cli.Context) {
				// $ notes new <note>
				editOrCreateNote(c.Args().First())
			},
		},
		{
			Name:      "edit",
			ShortName: "e",
			Usage:     "Edit a note",
			Action: func(c *cli.Context) {
				// $ notes edit <note>
				// $ notes e <note>
				editOrCreateNote(c.Args().First())
			},
		},
	}
	app.Run(os.Args)
}
