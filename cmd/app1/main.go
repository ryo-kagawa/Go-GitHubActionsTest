package main

import (
	"fmt"
	"os"
)

func main() {
	result, err := run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(result)
}

func run() (string, error) {
	args := GetArgs()
	if args.IsHelp() {
		return args.GetUsage(), nil
	}
	if args.IsVersion() {
		return args.GetVersion(), nil
	}
	return "app1 success", nil
}
