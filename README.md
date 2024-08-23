# pokedexcli
This is a command-line interface tool built in Golang that allows users to catch and retrieve information about Pokémon.

## Features
Every feature is mapped to its command line as below for use:
- `help`: Show help information for each commands.
- `exit`: Exit the program.
- `map`: List out location areas (20 areas for each page).
- `bmap`: Show the previous page of location areas.
- `explore {area name}`: Explore and get area information.
- `catch {pokemon name}`: Attempt to catch the Pokémon.
- `inspect {pokemon name}`: Get information of the caught pokemon.
- `pokedex`: List caught Pokémon.

## Prerequisites
Go 1.20 or higher installed on your machine.

## Run the program

### Clone the repository
```bash
git clone https://github.com/Xavier-Hsiao/pokedexcli.git
cd pokedexcli
```

### Build the CLI
```bash
go build pokedexcli
```

### Run the CLI
```bash
./pokedexcli
```

## Acknowledgments
- [PokeAPI](https://pokeapi.co/) - The Pokémon data is fetched from this fantastic API.