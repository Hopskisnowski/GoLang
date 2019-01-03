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
// type Response struct {
// 	Name    string    `json:"name"`
// 	Pokemon []Pokemon `json:"pokemon_entries"`
// }

// // A Pokemon Struct to map every pokemon to.
// type Pokemon struct {
// 	EntryNo int            `json:"entry_number"`
// 	Species PokemonSpecies `json:"pokemon_species"`
// }

// // A struct to map our Pokemon's Species which includes it's name
// type PokemonSpecies struct {
// 	Name string `json:"name"`
// }

type Response struct {
	Results []Results `json:"*"`
}

type Results struct {
	Ability string `json:"ability"`
	Berry   string `json:"berry"`
}

func main() {
	response, err := http.Get("https://pokeapi.co/api/v2/")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Results)

	for i := 0; i < len(responseObject.Results); i++ {
		fmt.Println(responseObject.Results[i].Ability + " " + responseObject.Results[i].Berry)
	}

}
