package main

import (
	"errors"
	"io"
	"os"
	"os/exec"
)

func OpenFileWithDefaultEditor(filename string) error {
	editorCommand := os.Getenv("EDITOR")
	if editCommand == emptyString {
		editCommand = os.Getenv("VISUAL")
		if editCommand == emptyString {
			return errors.New("you doesn't have set any editor, please add `EDITOR` or `VISUAL` to your configuration")
		}
	}
	cmd := exec.Command(editorCommand, filename)

	return cmd.Run()
}

func GetHostFilename() string {
	return "/etc/hosts"
}

func FprintNewLine(w io.Writer) (n int, err error) {
	newLine := []byte {10}
	return w.Write(newLine)
}
