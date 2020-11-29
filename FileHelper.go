package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func readFileInChunks(fn string) []Pokemon {
	var list []Pokemon
	f, err := os.Open(fn)
	if err != nil {
		log.Fatalf("Error to read [file=%v]: %v", fn, err.Error())
	}

	_, err1 := f.Stat()
	if err1 != nil {
		log.Fatalf("Could not obtain stat, handle error: %v", err.Error())
	}

	r := bufio.NewReader(f)
	d := json.NewDecoder(r)

	d.Token()
	for d.More() {
		pokemon := &Pokemon{}
		d.Decode(pokemon)
		list = append(list, *pokemon)
	}
	d.Token()
	return list
}

func readFile(fn string) string {
	content, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)
	return text
}

func createPokemonList(filename1 string, filename2 string, pokemonNumber int) {
	var pokemonsArray []string
	for i := 1; i <= pokemonNumber; i++ {
		var result = getPokemon(fmt.Sprint("https://pokeapi.co/api/v2/pokemon/", i))
		pokemonsArray = append(pokemonsArray, result)
	}
	savePokemons(pokemonsArray, filename1)
	savePokemons(reverse(pokemonsArray), filename2)
}

func savePokemons(pokemons []string, filename string) {
	f, err := os.Create(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	f.WriteString("[")
	var length = len(pokemons)
	for i := 0; i < length; i++ {
		f.WriteString(pokemons[i])
		if i != length-1 {
			f.WriteString(",")
		}
	}
	f.WriteString("]")
}
