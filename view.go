package main

import "github.com/urfave/cli"

func OnView(c *cli.Context) {
	OpenFileWithDefaultEditor(GetHostFilename())
}

var viewCommand = cli.Command{
	Name:    "view",
	Aliases: []string{"v"},
	Usage:   "View current hosts",
	Action:  OnView,
}
