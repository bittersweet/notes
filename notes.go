package main

import (
	"bufio"
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"os/exec"
	"path/filepath"
)

func showAllNotes() {
	matches, err := filepath.Glob("/Users/markmulder/dotfiles/notes/*.txt")
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

func showSubject(subject string) {
	path := fmt.Sprintf("/Users/markmulder/dotfiles/notes/%v.txt", subject)
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func getEditor() string {
	editor := os.Getenv("EDITOR")
	if len(editor) == 0 {
		editor = "vi"
	}
	return editor
}

func editNote(subject string) {
	file := fmt.Sprintf("/Users/markmulder/dotfiles/notes/%v.txt", subject)

	command := exec.Command(getEditor(), file)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Run()
}

func main() {
	app := cli.NewApp()
	app.Name = "notes"
	app.Usage = "Store your thoughts on all sorts of subjects"
	app.Action = func(c *cli.Context) {
		subject := c.Args().First()
		if len(subject) > 0 {
			showSubject(subject)
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
			Name:      "edit",
			ShortName: "e",
			Usage:     "Edit a note",
			Action: func(c *cli.Context) {
				editNote(c.Args().First())
			},
		},
	}
	app.Run(os.Args)

}
