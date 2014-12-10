package main

import (
  "fmt"
  "flag"
  "os"
  "bufio"
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

func main() {
  flag.Parse()
  subject := flag.Arg(0)
  if len(subject) == 0 {
    fmt.Printf("No subject given, exiting\n")
  } else {
    showSubject(subject)
  }
}
