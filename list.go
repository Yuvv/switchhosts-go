package main

import (
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
)

func OnList(c *cli.Context) error {
	appDir := AppDir()
	var fileNames []string
	var globalFileNames []string

	files, err := ioutil.ReadDir(appDir)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}
		isGlobal, hostFileName := GetConfigFilename(f.Name())
		if isGlobal {
			globalFileNames = append(globalFileNames, hostFileName)
		} else if hostFileName != emptyString {
			fileNames = append(fileNames, hostFileName)
		}
	}

	if c.GlobalBool(cliFlagGlobal) {
		fmt.Println("Global Config:")
		for _, ele := range globalFileNames {
			fmt.Print(ele, "\t")
		}
		fmt.Println()
	}

	fmt.Println("Normal Config:")
	for _, ele := range fileNames {
		fmt.Print(ele, "\t")
	}
	fmt.Println()
	return nil
}

var listCommand = cli.Command{
	Name:    "list",
	Aliases: []string{"l", "ls"},
	Usage:   "List existed hosts",
	Action:  OnList,
}
