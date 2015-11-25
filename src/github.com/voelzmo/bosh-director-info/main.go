package main

import (
	"fmt"
	"os"

	"github.com/voelzmo/bosh-director-info/director"
)

func main() {
	target, _ := parseArgs(os.Args)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	director := director.NewDirector(target)
	fmt.Printf("The director is targetting '%s'!\n", director.Status())
}

func parseArgs(args []string) (string, error) {
	expectedNumberOfArgs := 2
	if len(args) != expectedNumberOfArgs {
		return "", fmt.Errorf("parseArgs: Wrong number of arguments, expected %v, but got %v", expectedNumberOfArgs, len(args))
	}
	return args[1], nil
}
