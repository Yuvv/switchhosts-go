package main

import (
	"io"
	"os/exec"
)

func OpenFileWithDefaultEditor(filename string) error {
	cmd := exec.Command("notepad", filename)
	return cmd.Run()
}

func GetHostFilename() string {
	return "C:/WINDOWS/System32/drivers/etc/hosts"
}

func FprintNewLine(w io.Writer) (n int, err error) {
	newLine := []byte {13, 10}
	return w.Write(newLine)
}
