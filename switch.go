package main

import (
	"errors"
	"github.com/urfave/cli"
)

func OnSwitch(c *cli.Context) error {
	global := c.GlobalBool(cliFlagGlobal)
	if !c.Args().Present() {
		return errors.New("env cannot be null")
	}

	configMap := map[string]bool{}
	for _, ele := range c.Args() {
		configMap[ele] = true
	}
	configArray := make([]string, len(configMap))
	i := 0
	for key := range configMap {
		configArray[i] = key
		i++
	}
	return SwitchConfig(global, configArray...)
}

var switchCommand = cli.Command{
	Name:    "switch",
	Aliases: []string{"s", "sw"},
	Usage:   "Switch to hosts",
	Action:  OnSwitch,
}
