package main

import (
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

func SwitchConfig(global bool, configs ...string) error {
	configPaths := make([]string, len(configs))
	for i, ele := range configs {
		if strings.Contains(ele, ".") {
			return errors.New("config name cannot contains \".\"")
		}
		configPaths[i] = GetConfigFullPath(ele, global)
	}

	hostsFile, err := os.OpenFile(GetHostFilename(), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if os.IsExist(err) {
		return errors.New("open hosts file failed\n" + err.Error())
	}
	defer hostsFile.Close()

	// add head
	hostsFile.WriteString("# ****************** SwitchHosts-go ******************")
	FprintNewLine(hostsFile)
	FprintNewLine(hostsFile)

	// add global config at first
	appDir := AppDir()
	files, err := ioutil.ReadDir(appDir)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}
		if IsGlobalHostFile(f.Name()) {
			file, err := os.Open(path.Join(appDir, f.Name()))
			if os.IsExist(err) {
				return errors.New("write " + path.Ext(f.Name()) + " failed")
			}

			io.Copy(hostsFile, file)
			hostsFile.WriteString("\n\n")
			file.Close()
		}
	}
	// and then add normal configs
	for _, filename := range configPaths {
		if !PathExist(filename) {
			fmt.Printf("Config file %s not exist, ignored\n", path.Ext(filename))
			continue
		}

		file, err := os.Open(filename)
		if os.IsExist(err) {
			return errors.New("write " + path.Ext(filename) + " failed")
		}

		io.Copy(hostsFile, file)
		hostsFile.WriteString("\n\n")
		file.Close()
	}

	return nil
}

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
	if global {
		log.Println("global flag is ignored when switch configs")
	}
	return SwitchConfig(global, configArray...)
}

var switchCommand = cli.Command{
	Name:    "switch",
	Aliases: []string{"s", "sw"},
	Usage:   "Switch to hosts",
	Action:  OnSwitch,
}
