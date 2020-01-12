package main

import "os/exec"

func OpenFileWithDefaultEditor(filename string) error {
	cmd := exec.Command("open", "-e", filename)
	return cmd.Run()
}

func GetHostFilename() string {
	return "/etc/hosts"
}

func FprintNewLine(w io.Writer) (n int, err error) {
	newLine := []byte {10}
	return w.Write(newLine)
}
