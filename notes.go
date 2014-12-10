package main

import (
  "fmt"
  "flag"
  "os"
  "bufio"
)

func showSubject(subject string) {
  file, err := os.Open("/Users/markmulder/dotfiles/notes/cli.txt")
  if err != nil {
    fmt.Println("error in showSubject")
    fmt.Printf("%v\n", err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    fmt.Println(scanner.Text())
  }
}

func main() {
  flag.Parse()
  subject := flag.Arg(0)
  if len(subject) == 0 {
    fmt.Printf("No subject given, exiting\n")
  } else {
    showSubject(subject)
  }
}
