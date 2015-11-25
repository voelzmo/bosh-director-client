package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/voelzmo/bosh-director-info/director"
)

func main() {
	target, rootCAPath, err := parseArgs(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	director := director.NewDirector(target, rootCAPath)
	prettyStatus, _ := json.MarshalIndent(director.Status(), "", "  ")

	fmt.Printf("The director: '%s'!\n", prettyStatus)
}

func parseArgs(args []string) (string, string, error) {
	expectedNumberOfArgs := 3
	if len(args) != expectedNumberOfArgs {
		return "", "", fmt.Errorf("parseArgs: Wrong number of arguments, expected %v, but got %v", expectedNumberOfArgs, len(args))
	}
	return args[1], args[2], nil
}
