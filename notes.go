package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	ListPtr bool
)

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

func setOptions() {
	flag.BoolVar(&ListPtr, "list", false, "list all known notes")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s subject [options]\n", os.Args[0])
		flag.PrintDefaults()
	}
}

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

func main() {
	setOptions()
	flag.Parse()

	subject := flag.Arg(0)

	if ListPtr || len(subject) == 0 {
		showAllNotes()
		os.Exit(0)
	}

	showSubject(subject)
}
