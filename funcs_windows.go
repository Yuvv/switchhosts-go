package main

import (
	"os/exec"
)

func OpenFileWithDefaultEditor(filename string) error {
	cmd := exec.Command("notepad", filename)
	return cmd.Run()
}


func GetHostFilename() string {
	return "C:/WINDOWS/System32/drivers/etc/hosts"
}