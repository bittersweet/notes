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

type Note struct {
	Explanation []string
	Command     []string
}

func (n *Note) Print() {
	for _, line := range n.Explanation {
		colorizeComment(line)
	}

	for _, line := range n.Command {
		fmt.Println(line)
	}
	fmt.Println()
}

func FindNotes(notes []Note, query string) []Note {
	var results []Note

Loop:
	for _, note := range notes {
		queryRegexp := fmt.Sprintf("(?i)%s", query)
		searchRegexp, _ := regexp.Compile(queryRegexp)

		for _, line := range note.Explanation {
			if searchRegexp.MatchString(line) == true {
				results = append(results, note)
				continue Loop
			}
		}

		for _, line := range note.Command {
			if searchRegexp.MatchString(line) == true {
				results = append(results, note)
				continue Loop
			}
		}
	}

	return results
}

func showAllNotes() {
	path := fmt.Sprintf("%v*.txt", getNotesDir())
	matches, err := filepath.Glob(path)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	for i := 0; i < len(matches); i++ {
		match := matches[i]
		filename := filepath.Base(match)
		base_size := len(filename) - 4 // Note filename without .txt
		fmt.Println(filename[:base_size])
	}
}

func showNote(params ...string) {
	note := params[0]
	var query string
	if len(params) > 1 {
		query = params[1]
	}

	path := fmt.Sprintf("%v%v.txt", getNotesDir(), note)
	file, err := os.Open(path)
	if err != nil {
		// File does not exist, creating it
		editOrCreateNote(note)
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
		notes = FindNotes(notes, query)
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
	file := fmt.Sprintf("%v%v.txt", getNotesDir(), note)

	command := exec.Command(getEditor(), file)
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

func getNotesDir() string {
	homedir := os.Getenv("HOME")
	notesdir := os.Getenv("NOTESDIR")
	if notesdir != "" {
		return notesdir
	} else {
		return fmt.Sprintf("%v/dotfiles/notes/", homedir)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "notes"
	app.Version = "0.5.0"
	app.Usage = "Store your thoughts on all sorts of subjects"
	app.Action = func(c *cli.Context) {
		note := c.Args().First()
		if len(c.Args()) == 0 {
			showAllNotes()
		} else if len(c.Args()) == 1 {
			showNote(note)
		} else {
			showNote(note, c.Args()[1])
		}
	}
	app.Commands = []cli.Command{
		{
			Name:      "list",
			ShortName: "l",
			Usage:     "List all notes",
			Action: func(c *cli.Context) {
				showAllNotes()
			},
		},
		{
			Name:      "new",
			ShortName: "n",
			Usage:     "Create new note",
			Action: func(c *cli.Context) {
				editOrCreateNote(c.Args().First())
			},
		},
		{
			Name:      "edit",
			ShortName: "e",
			Usage:     "Edit a note",
			Action: func(c *cli.Context) {
				editOrCreateNote(c.Args().First())
			},
		},
	}
	app.Run(os.Args)
}
