package main

import (
	"errors"
	"github.com/urfave/cli"
)

func OnSwitch(c *cli.Context) error {
	global := c.GlobalBool(cliFlagGlobal)
	envName := c.Args().First()
	if envName == emptyString {
		return errors.New("env cannot be null")
	}
	// todo
	return nil
}

var switchCommand = cli.Command{
	Name:    "switch",
	Aliases: []string{"s", "sw"},
	Usage:   "Switch to hosts",
	Action:  OnSwitch,
}