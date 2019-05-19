package main

import "github.com/urfave/cli"

func OnEdit(c *cli.Context) error {
	// todo:
	return nil
}

var editCommand = cli.Command{
	Name:    "edit",
	Aliases: []string{"e"},
	Usage:   "Edit existed hosts",
	Action:  OnEdit,
}
