package main

import "strings"

func cleanInput(text string) []string {
	// split the user's input into words based on whitespace
	// lowercase the input and trim any leading or trailing whitespaces
	words := strings.Fields(strings.ToLower(text))
	return words
}
