# notes

[![Build Status](https://drone.io/github.com/bittersweet/notes/status.png)](https://drone.io/github.com/bittersweet/notes/latest)

Store your thoughts on all sorts of subjects and easily read them on the
commandline. I use it for CLI commands I don't use very often and keep
forgetting.

It expects a `~/dotfiles/notes/` directory to store all the notes but you can
override this by setting the `$NOTESDIR` variable from within your `.profile`
for example.
The notes are stored as simple textfiles for easy editting and portability.

It works best if you organize your notes in the following format:

```
# Show open files by PID
lsof -p $PID

# pretty print JSON file
cat file.json | python -m json.tool > pretty.json

```

By using a newline, notes will now its a new command, which is handy when
searching within notes by doing `notes subject query`.
The explanation line, prepended with #, will make sure this is colorized in
your terminal, the previous example will look something like the following:

![](http://img.springe.st/1._tmux_2015-06-12_17-15-19.png)

## Installation

You can download the [latest
release](https://github.com/bittersweet/notes/releases) or build it yourself of
course.

## Usage

* `notes` Lists all the notes you have available
* `notes [note]` show the selected note
* `notes [note] [query]` shows only matching notes within the `[note]` file
* `notes edit [note]` edits
* `notes new [name]` creates a new note

Editting and creating new notes uses `$EDITOR` or `vi` in case that is not set.

## Testing

To run the tests simply use the `go test` command.
