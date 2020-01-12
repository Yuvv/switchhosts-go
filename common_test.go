package main

import "testing"

func TestPathExist(t *testing.T) {
	fakePath := "./abc/a"
	exists := PathExist(fakePath)
	if exists {
		t.Errorf("path `%s` should be not exists but it is", fakePath)
	}

	curPath := "."
	exists = PathExist(curPath)
	if !exists {
		t.Errorf("path `%s` should be exists but it is not", fakePath)
	}
}

func TestIsGlobalHostFile(t *testing.T) {
	configName1 := _globalHostPrefix + "abc"
	if !IsGlobalHostFile(configName1) {
		t.Errorf("`%s` expected to be a global config, but normal returned", configName1)
	}

	configName2 := _normalHostPrefix + "abc"
	if IsGlobalHostFile(configName2) {
		t.Errorf("`%s` expected to be a normal config, but global returned", configName2)
	}
}

func TestIsNormalHostFile(t *testing.T) {
	configName1 := _globalHostPrefix + "abc"
	if IsNormalHostFile(configName1) {
		t.Errorf("`%s` expected to be a global config, but normal returned", configName1)
	}

	configName2 := _normalHostPrefix + "abc"
	if !IsNormalHostFile(configName2) {
		t.Errorf("`%s` expected to be a normal config, but global returned", configName2)
	}
}