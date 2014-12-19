# notes

Store your thoughts on all sorts of subjects and easily read them on the
commandline. I use it for CLI commands I don't use very often and keep
forgetting.

It expects a `~/dotfiles/notes/` directory to store all the notes but you can
override this by setting the `$NOTESDIR` variable from within your `.profile`
for example.
The notes are stored as simple textfiles for easy editting and portability.

You can prepend a line with `#` to explain the command, these comments will get
colorized, an example:

```
# Show open files by PID
lsof -p $PID
```

![](http://img.springe.st/1._tmux_2014-12-11_11-23-42.png)

## Installation

You can download the [latest
release](https://github.com/bittersweet/notes/releases) or build it yourself of
course.

## Usage

* `notes` Lists all the notes you have available
* `notes [note]` show the selected note
* `notes edit [note]` edits
* `notes new [name]` creates a new note

Editting and creating new notes uses `$EDITOR` or `vi` in case that is not set.

## Testing

To run the tests simply use the `go test` command.
