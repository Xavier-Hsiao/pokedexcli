package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Get user input
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter your input")

	scanner.Scan()
	input := scanner.Text()

	fmt.Println("You've entered:", input)
}
