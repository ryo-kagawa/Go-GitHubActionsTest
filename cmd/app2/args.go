package main

import (
	"os"
	"strings"
)

var usage = `
Usage: app
Usage: app --help
    ここを表示します
Usage: app --version
    バージョンを表示します
`

type Args struct {
	help    bool
	version bool
}

func GetArgs() Args {
	result := Args{}
	args := os.Args[1:]
	if 0 == len(args) {
		return result
	}
	switch args[0] {
	case "--help":
		return Args{
			help: true,
		}
	case "--version":
		return Args{
			version: true,
		}
	}
	return Args{}
}

func (_self Args) IsHelp() bool {
	return _self.help
}

func (_self Args) IsVersion() bool {
	return _self.version
}

func (_self Args) GetUsage() string {
	return strings.Trim(usage, "\n")
}

func (_self Args) GetVersion() string {
	return GetVersion()
}
