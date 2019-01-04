package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// URL to code https://tutorialedge.net/golang/consuming-restful-api-with-go/

// // A Response struct to map the Entire Response
type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

// // A Pokemon Struct to map every pokemon to.
type Pokemon struct {
	EntryNo byte           `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

// // A struct to map our Pokemon's Species which includes it's name
type PokemonSpecies struct {
	Name string `json:"name"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	stringHolder := []string{string(responseData)}

	for i := 0; i < len(stringHolder); i++ {
		//fmt.Println(stringHolder[i])
	}

	//Printing content of stringHolder, which is only stringHolder[0] eventhough it is slice strings
	fmt.Println(stringHolder)

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	//fmt.Println(responseObject.Name)

	// for i := 0; i < len(responseObject.Results); i++ {
	// 	fmt.Println(responseObject.Results[i].Ability + " " + responseObject.Results[i].Berry)
	// }

	f, err := os.Create("/tmp/dat2")
	check(err)

	defer f.Close()

	for i := 0; i < len(responseObject.Pokemon); i++ {
		if responseObject.Pokemon[i].Species.Name == "kabuto" {
			fmt.Printf("%s fundet med entry nummer %d \n", responseObject.Pokemon[i].Species.Name, responseObject.Pokemon[i].EntryNo)
			f.WriteString(responseObject.Pokemon[i].Species.Name)
			//f.Write(string(responseObject.Pokemon[i].EntryNo))
		}
		// else {
		// 	fmt.Printf("%s Fundet \n", responseObject.Pokemon[i].Species.Name)
		// }
	}

}
