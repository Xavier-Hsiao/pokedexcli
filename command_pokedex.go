package main

import "fmt"

func commandPokedex(cfg *config, args []string) error {
	pokedex := cfg.pokedex

	fmt.Println("Your pokedex:")
	for name, _ := range pokedex {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
