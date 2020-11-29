package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

const POKEMON_NUMBER = 500

func main() {

	var pokemonName1 = "pokemons1.json"
	var pokemonName2 = "pokemons2.json"
	createPokemonList(pokemonName1, pokemonName2, POKEMON_NUMBER)

	readPokemonsAllIntoMemory(pokemonName1, pokemonName2)
	var pokemonList1, pokemonList2 []Pokemon
	pokemonList1, pokemonList2 = readPokemonsInMemoryOptimisedWay(pokemonName1, pokemonName2)

	compareTwoJson(pokemonList1, pokemonList2)
	compareTwoJsonWithMap(pokemonList1, pokemonList2)
	compareTwoJsonWithMapParallel(pokemonList1, pokemonList2)

}

func readPokemonsAllIntoMemory(pokemonName1 string, pokemonName2 string) ([]Pokemon, []Pokemon) {
	start := time.Now()
	var pokemonList1 = readFile(pokemonName1)
	var pokemonList1Json []Pokemon

	err := json.Unmarshal([]byte(pokemonList1), &pokemonList1Json)
	if err != nil {
		log.Println(err)
	}

	var pokemonList2 = readFile(pokemonName2)
	var pokemonList2Json []Pokemon

	err2 := json.Unmarshal([]byte(pokemonList2), &pokemonList2Json)
	if err2 != nil {
		log.Println(err2)
	}
	elapsed := time.Since(start)
	fmt.Println("Reading Pokemons took ", elapsed)
	return pokemonList1Json, pokemonList2Json
}

func readPokemonsInMemoryOptimisedWay(pokemonName1 string, pokemonName2 string) ([]Pokemon, []Pokemon) {
	start := time.Now()
	var pokemonList1 = readFileInChunks(pokemonName1)
	var pokemonList2 = readFileInChunks(pokemonName2)
	elapsed := time.Since(start)
	fmt.Println("Reading Pokemons Optimised took ", elapsed)
	return pokemonList1, pokemonList2
}

func compareTwoJsonWithMapParallel(pokemonList1Json []Pokemon, pokemonList2Json []Pokemon) {
	fmt.Println("========================")
	fmt.Println("Parallel Compare Started")
	EqualWithMapParallel(pokemonList1Json, pokemonList2Json)
	fmt.Println("Parallel Compare Finished")
	fmt.Println("========================")
}

func compareTwoJsonWithMap(pokemonList1Json []Pokemon, pokemonList2Json []Pokemon) {

	fmt.Println("========================")
	fmt.Println("Map Compare Started")
	EqualWithMap(pokemonList1Json, pokemonList2Json)
	fmt.Println("Map Compare Finished")
	fmt.Println("========================")
}

func compareTwoJson(pokemonList1Json []Pokemon, pokemonList2Json []Pokemon) {

	fmt.Println("========================")
	fmt.Println("Simple Compare Started")
	Equal(pokemonList1Json, pokemonList2Json)
	fmt.Println("Simple Compare Finished")
	fmt.Println("========================")
}
