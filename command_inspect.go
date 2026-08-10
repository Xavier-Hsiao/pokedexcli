package main

import "fmt"

func commandInspect(cfg *config, args []string) error {
	pokedex := cfg.pokedex
	pokemon, ok := pokedex[args[0]]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Printf(" - %s\n", typeInfo.Type.Name)
	}

	return nil
}
