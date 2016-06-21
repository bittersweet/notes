package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/urfave/cli"
)

// Note holds parsed explanation and commands
type Note struct {
	Explanation []string
	Command     []string
}

// Print outputs an indivial note explanation and command to STDOUT.
func (n *Note) Print() {
	for _, line := range n.Explanation {
		colorizeComment(line)
	}

	for _, line := range n.Command {
		fmt.Println(line)
	}
}

// hasData checks if there is an explanation and command set for this note.
func (n *Note) hasData() bool {
	if len(n.Explanation) > 0 && len(n.Command) > 0 {
		return true
	}

	return false
}

// findNotes looks for <query> in all notes passed in, returns matching notes.
func findNotes(notes []Note, query string) []Note {
	var results []Note

Loop:
	for _, note := range notes {
		queryRegexp := fmt.Sprintf("(?i)%s", query)
		searchRegexp, _ := regexp.Compile(queryRegexp)

		for _, line := range note.Explanation {
			if searchRegexp.MatchString(line) == true {
				results = append(results, note)
				// We already found a match, continue with the next note
				continue Loop
			}
		}

		for _, line := range note.Command {
			if searchRegexp.MatchString(line) == true {
				results = append(results, note)
				// We already found a match, continue with the next note
				continue Loop
			}
		}
	}

	return results
}

// searchAllNotes looks through all note files and prints matching notes.
func searchAllNotes(query string) {
	filenames := getAllNotes()

	var matchedNotes []Note
	for _, filename := range filenames {
		// parse filename and get a []Note
		note, err := parseNoteFile(filename)
		if err != nil {
			log.Fatal("note did not exist")
		}

		// check if the notes contain our query
		matches := findNotes(note, query)
		for _, match := range matches {
			matchedNotes = append(matchedNotes, match)
		}
	}

	for _, note := range matchedNotes {
		note.Print()
	}

	// Add empty newline
	fmt.Println()
}

// getAllNotes returns all .txt files in the notes dir, with the extension
// removed.
func getAllNotes() []string {
	path := fmt.Sprintf("%v/*.txt", getNotesDir())
	matches, err := filepath.Glob(path)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	for i := 0; i < len(matches); i++ {
		match := matches[i]
		filename := filepath.Base(match)
		baseSize := len(filename) - 4 // Strip .txt from note filenames
		matches[i] = filename[:baseSize]
	}

	return matches
}

// showAllNotes prints out all note filenames.
func showAllNotes() {
	notes := getAllNotes()

	for _, note := range notes {
		fmt.Println(note)
	}
}

// parseNoteFile reads a file from `note` and converts that to a slice of
// `Note`s. It will return an err if the file can not be opened (for now I'm
// assuming that always means it does not exist yet).
func parseNoteFile(note string) ([]Note, error) {
	path := getNote(note)
	file, err := os.Open(path)
	if err != nil {
		// File does not exist
		return nil, err
	}
	defer file.Close()

	var notes []Note
	var n Note

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	commentRegexp, _ := regexp.Compile("^#")

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		var previousLine string
		if i == 0 {
			previousLine = ""
		} else {
			previousLine = lines[i-1]
		}

		if commentRegexp.MatchString(line) == true {
			// We found a # with the preceding string being blank, which means the
			// start of a new note.
			if previousLine == "" {
				// Only add it when it contains an explanation and command, so we don't
				// add a blank note.
				if n.hasData() {
					notes = append(notes, n)
				}
				n = Note{}
			}
			n.Explanation = append(n.Explanation, line)
		} else {
			n.Command = append(n.Command, line)
		}
	}

	// We are done going over all the lines. If the file ended without a blank
	// line it was not appended yet in the loop, so do that here.
	if n.hasData() {
		notes = append(notes, n)
	}

	return notes, nil
}

// showNote prints out a complete note file with all notes, adding a newline in
// between and leveraging colorizeComment to pretty print the explanation.
func showNote(params ...string) {
	note := params[0]

	notes, err := parseNoteFile(note)
	if err != nil {
		editOrCreateNote(note)
	}

	// If we received multiple parameters, the second one will be the query, we
	// filter our parsed notes here.
	if len(params) > 1 {
		query := params[1]
		notes = findNotes(notes, query)
	}

	for _, note := range notes {
		note.Print()
	}

	// Add empty newline
	fmt.Println()
}

// colorizeComment pretty prints a note, the comment description gets a
// different color.
func colorizeComment(line string) {
	highlight := "\033[33m"
	reset := "\033[0m"

	fmt.Printf("%v%v%v\n", highlight, line, reset)
}

// editOrCreateNote opens <note> in your editor.
func editOrCreateNote(note string) {
	path := getNote(note)

	command := exec.Command(getEditor(), path)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Run()
}

// getEditor checks your environment variables to see which editor you use,
// falling back to vi.
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
	app.Version = "0.6.2"
	app.Usage = "Store your thoughts on all sorts of subjects"
	app.EnableBashCompletion = true
	app.BashComplete = func(c *cli.Context) {
		notes := getAllNotes()
		for _, note := range notes {
			fmt.Println(note)
		}
	}
	app.Action = func(c *cli.Context) error {
		note := c.Args().First()
		if len(c.Args()) == 0 {
			// We received no arguments, equal to:
			// $ notes
			showAllNotes()
		} else if len(c.Args()) == 1 {
			// We received 1 argument, to open a specific note, equal to:
			// $ notes <note>
			showNote(note)
		} else {
			// We received 2 arguments, to open a specific note and search, equal to:
			// $ notes <note> <query>
			showNote(note, c.Args()[1])
		}
		return nil
	}
	app.Commands = []cli.Command{
		{
			Name:      "list",
			ShortName: "l",
			Usage:     "List all notes",
			Action: func(c *cli.Context) error {
				// $ notes list
				showAllNotes()
				return nil
			},
		},
		{
			Name:      "new",
			ShortName: "n",
			Usage:     "Create new note",
			Action: func(c *cli.Context) error {
				// $ notes new <note>
				editOrCreateNote(c.Args().First())
				return nil
			},
		},
		{
			Name:      "edit",
			ShortName: "e",
			Usage:     "Edit a note",
			Action: func(c *cli.Context) error {
				// $ notes edit <note>
				// $ notes e <note>
				editOrCreateNote(c.Args().First())
				return nil
			},
		},
		{
			Name:      "search",
			ShortName: "s",
			Usage:     "Search through all your notes",
			Action: func(c *cli.Context) error {
				// $ notes search <query>
				// $ notes s <query>
				query := c.Args().First()
				searchAllNotes(query)
				return nil
			},
		},
	}
	app.Run(os.Args)
}
