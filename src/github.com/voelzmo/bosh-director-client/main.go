package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/voelzmo/bosh-director-client/director"
)

func main() {
	target, rootCAPath, clientName, clientSecret, err := parseArgs(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	director := director.NewDirector(target, rootCAPath, clientName, clientSecret)
	prettyStatus, _ := json.MarshalIndent(director.Status(), "", "  ")

	fmt.Printf("The director status: '%s'\n", prettyStatus)

	prettyLogin, _ := json.MarshalIndent(director.Login(), "", "  ")

	fmt.Printf("The director login: '%s'\n", prettyLogin)

}

func parseArgs(args []string) (string, string, string, string, error) {
	expectedNumberOfArgs := 5
	if len(args) != expectedNumberOfArgs {
		return "", "", "", "", fmt.Errorf("parseArgs: Wrong number of arguments, expected %v, but got %v\nUsage: bosh-director-client <director URL> <root CA path> <oauth client name> <oauth client secret>", expectedNumberOfArgs-1, len(args)-1)
	}
	return args[1], args[2], args[3], args[4], nil
}
