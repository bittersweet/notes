# Notes Changelog

## 0.7.2

- Nix compatibility via `make install`.

## 0.7.1

- Project now uses Go modules.

## 0.7.0

- Added searching within all notes with `search` or `s`. Use `$ notes search
xargs` to print notes from all files that match.
- Refactored note parsing so you can have newlines within 1 note, thanks
@mongrelion for [the bugreport](https://github.com/bittersweet/notes/issues/4).

## 0.6.2

- Fixed opening notes that broke in 0.6.1 because of the removal of the trailing slash :-).

## 0.6.1

- Fixed using a custom `NOTESDIR` not working without a trailing slash.

## 0.6.0

- Added searching in notes via `notes subject query` to only display matched
notes, searches on the explanation and command.

## 0.5.0

- Allow overriding of storage directory via `NOTESDIR`, thanks to @mongrelion.

## 0.4.0

- Opens editor with note name anyway if it doesn't exist yet.

## 0.3.0

- Colorize commented lines.

## 0.2.0

- Add creating new notes.

## 0.1.0

- Initial release.
- Listing, showing and editing notes is available.
