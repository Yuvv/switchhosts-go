package main

import (
	"fmt"
	"github.com/urfave/cli"
)

func OnDelete(c *cli.Context) error {
	global := c.GlobalBool(cliFlagGlobal)
	envName := c.Args().First()
	if envName == emptyString {
		fmt.Println("env cannot be null")
		return nil
	}
	return DelHostFileByName(envName, global)
}

var delCommand = cli.Command{
	Name:    "del",
	Aliases: []string{"d"},
	Usage:   "Delete new hosts",
	Action:  OnDelete,
}
