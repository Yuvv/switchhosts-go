package main

import (
	"log"

	"github.com/urfave/cli"
)

func OnView(c *cli.Context) {
	if err := OpenFileWithDefaultEditor(GetHostFilename()); err != nil {
		log.Fatalf("cannot open file, error:%+v", err)
	}
}

var viewCommand = cli.Command{
	Name:    "view",
	Aliases: []string{"v"},
	Usage:   "View current hosts",
	Action:  OnView,
}
