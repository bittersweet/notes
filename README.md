# notes

Store your thoughts on all sorts of subjects and easily read them on the
commandline. I use it for CLI commands I don't use very often and keep
forgetting.

## Installation

As it's in super early stage, you'll have to build it yourself. I'm hoping you
have Go already setup. If so, just clone, `go get`, and run `make` to build it
and move it to `/usr/local/bin/`, or `go build notes.go` and move it yourself to
a desired location.

It also expects a `~/dotfiles/notes/` directory to store all the notes. These
are just simple textfiles for easy editting and portability.

## Usage

`notes` Lists all the notes you have available
`notes [note]` show the selected note
`notes edit [note]` edits
`notes new [name]` creates a new note

Editting and creating new notes uses $EDITOR or vi if that is not set.

