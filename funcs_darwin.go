package main

import "os/exec"

func OpenFileWithDefaultEditor(filename string) error {
	cmd := exec.Command("open", "-e", filename)
	return cmd.Run()
}


func GetHostFilename() string {
	return "/etc/hosts"
}