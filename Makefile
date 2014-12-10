all:
	@@go build notes.go
	@@cp notes /usr/local/bin/
	@@echo "built and moved"
