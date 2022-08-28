package main

import (
	"errors"
	"io"
	"log"
	"os"
	"os/exec"
)

var editorCandidates = []string{
	"/usr/bin/vim",
	"/usr/bin/vi",
	"/bin/nano",
	"/usr/bin/nano",
}

func OpenFileWithDefaultEditor(filename string) error {
	editorCommand := os.Getenv("EDITOR")
	if editorCommand == emptyString {
		editorCommand = os.Getenv("VISUAL")
		if editorCommand == emptyString {
			for _, candidate := range editorCandidates {
				if fInfo, err := os.Stat(candidate); err == nil && (fInfo.Mode()&0x111) > 0 {
					log.Printf("`EDITOR`/`VISUAL` not set, use backup editor: `%s`\n", candidate)
					editorCommand = candidate
					break
				}
			}
		}
		if editorCommand == emptyString {
			return errors.New("you doesn't have any editor, please add `EDITOR` or `VISUAL` to your configuration or install vim/neovim/nano")
		}
	}
	cmd := exec.Command(editorCommand, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func GetHostFilename() string {
	return "/etc/hosts"
}

func FprintNewLine(w io.Writer) (n int, err error) {
	newLine := []byte{10}
	return w.Write(newLine)
}
