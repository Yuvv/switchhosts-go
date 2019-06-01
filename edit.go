package main

import (
	"errors"
	"github.com/urfave/cli"
)

func OnEdit(c *cli.Context) error {
	global := c.GlobalBool(cliFlagGlobal)
	envName := c.Args().First()
	if envName == emptyString {
		return errors.New("env cannot be null")
	}

	configFullPath := GetConfigFullPath(envName, global)

	return OpenFileWithDefaultEditor(configFullPath)
}

var editCommand = cli.Command{
	Name:    "edit",
	Aliases: []string{"e"},
	Usage:   "Edit existed hosts",
	Action:  OnEdit,
}
