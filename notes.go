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
	path := fmt.Sprintf("%v*.txt", getHomeDir())
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
	path := fmt.Sprintf("%v%v.txt", getHomeDir(), note)
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

func editNote(note string) {
	file := fmt.Sprintf("%v%v.txt", getHomeDir(), note)

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

func getHomeDir() string {
	home := os.Getenv("HOME")
	return fmt.Sprintf("%v/dotfiles/notes/", home)
}

func main() {
	app := cli.NewApp()
	app.Name = "notes"
	app.Version = "0.1.0"
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
