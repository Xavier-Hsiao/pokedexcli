package main

func main() {
	startRpl()

	// We have two commands:
	// help: prints a help message describing how to use the REPL
	// exit: exits the program
	type cliCommand struct {
		name        string
		description string
	}
}
