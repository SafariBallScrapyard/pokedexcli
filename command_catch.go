package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a Pokémon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.CatchPokemon(name)
	if err != nil {
		return err
	}

	num := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokéball at %s...\n", pokemon.Name)
	if num > 35 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	cfg.caughtPokemon[pokemon.Name] = pokemon
	fmt.Println("You may now inspect it with the 'inspect' command.")

	return nil
}
