package main

import (
	"errors"
	"os"
	"os/exec"
	"github.com/urfave/cli"
)

func OnAdd(c *cli.Context) error {
	if !c.Args().Present() {
		return errors.New("Name of profile needed")
	}

	configName := c.Args().First()
	global := c.GlobalBool(cliFlagGlobal)

	if IsConfigExist(configName, global) {
		return errors.New("Config file exists")
	}

	configFullPath, err := AddConfig(configName, global)
	if os.IsExist(err) {
		return err
	}
	
	editor := exec.Command("open", "-e", configFullPath)
	err = editor.Run()
	return err
}

var addCommand = cli.Command{
	Name:    "add",
	Aliases: []string{"a"},
	Usage:   "Add new hosts",
	Action:  OnAdd,
}
