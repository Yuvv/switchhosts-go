package main

import (
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"os"
	"strings"
)

func AddConfig(name string, global bool) (string, error) {
	fullPath := GetConfigFullPath(name, global)

	file, err := os.Create(fullPath)
	if os.IsExist(err) {
		return emptyString, errors.New("create failed")
	}
	fmt.Fprintf(file, "# SwitchHosts-Go =======> %s", name)
	fmt.Fprintln(file)
	file.Close()

	return fullPath, nil
}

func OnAdd(c *cli.Context) error {
	if !c.Args().Present() {
		return errors.New("name of profile needed")
	}

	configName := c.Args().First()
	global := c.GlobalBool(cliFlagGlobal)

	if strings.Contains(configName, ".") {
		return errors.New("config name cannot contains \".\"")
	}

	if IsConfigExist(configName, global) {
		return errors.New("config file existed")
	}

	configFullPath, err := AddConfig(configName, global)
	if os.IsExist(err) {
		return err
	}

	err = OpenFileWithDefaultEditor(configFullPath)
	return err
}

var addCommand = cli.Command{
	Name:    "add",
	Aliases: []string{"a"},
	Usage:   "Add new hosts",
	Action:  OnAdd,
}
