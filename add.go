package main

import (
	"fmt"
	"github.com/urfave/cli"
)

func OnAdd(c *cli.Context) error {

	fmt.Println(c.Args().First())

	// todo:
	return nil
}

var addCommand = cli.Command{
	Name:    "add",
	Aliases: []string{"a"},
	Usage:   "Add new hosts",
	Action:  OnAdd,
}
