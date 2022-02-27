package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/ryo-kagawa/go-utils/conditional"
)

func main() {
	result, err := run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(
		conditional.StringFunc(
			2 <= len(os.Args),
			func() string {
				return os.Args[1]
			},
			func() string {
				return ""
			},
		),
	)
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
	return "", errors.New("app2 error")
}
