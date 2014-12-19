package main

import (
	"bufio"
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

func showAllNotes() {
	path := fmt.Sprintf("%v*.txt", getNotesDir())
	matches, err := filepath.Glob(path)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	for i := 0; i < len(matches); i++ {
		match := matches[i]
		filename := filepath.Base(match)
		base_size := len(filename) - 4
		fmt.Println(filename[:base_size])
	}
}

func showNote(note string) {
	path := fmt.Sprintf("%v%v.txt", getNotesDir(), note)
	file, err := os.Open(path)
	if err != nil {
		// File does not exist, creating it
		editOrCreateNote(note)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		commentRegexp, _ := regexp.Compile("^#")
		if commentRegexp.MatchString(line) == true {
			colorizeComment(line)
		} else {
			fmt.Println(line)
		}
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
		if len(note) > 0 {
			showNote(note)
		} else {
			showAllNotes()
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
