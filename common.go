package main

import (
	"os"
	"os/user"
	"path"
	"strings"
)

const _configDir = ".config"
const _appDir = "SwitchHosts-Go"
const _globalHostPrefix = "g."
const _normalHostPrefix = "n."

const emptyString = ""
const cliFlagGlobal = "global"

func AppDir() string {
	curUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	userHomeDir := curUser.HomeDir
	appDir := path.Join(userHomeDir, _configDir, _appDir)
	if _, err := os.Stat(_configDir); os.IsNotExist(err) {
		err = os.MkdirAll(appDir, os.ModeDir)
		if os.IsExist(err) {
			panic(err)
		}
	}
	return appDir
}

func IsGlobalHostFile(filename string) bool {
	return strings.HasPrefix(filename, _globalHostPrefix)
}

func IsNormalHostFile(filename string) bool {
	return strings.HasPrefix(filename, _normalHostPrefix)
}

func GetHostFileName(filename string) (bool, string) {
	if IsGlobalHostFile(filename) {
		return true, strings.Replace(filename, _globalHostPrefix, emptyString, 1)
	} else if IsNormalHostFile(filename) {
		return false, strings.Replace(filename, _normalHostPrefix, emptyString, 1)
	}
	return false, emptyString
}

func DelHostFileByName(name string, global bool) error {
	appDir := AppDir()
	var fullPath string
	if global {
		fullPath = path.Join(appDir, _globalHostPrefix+name)
	} else {
		fullPath = path.Join(appDir, _normalHostPrefix+name)
	}
	return os.Remove(fullPath)
}
