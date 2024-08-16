package main

import "fmt"

func callHelp(cfg *config, args ...string) error {
	fmt.Println("Here are the available commands:")

	availableCommands := getCommands()
	for cmd := range availableCommands {
		fmt.Printf("- %s\n", cmd)
	}
	return nil
}
