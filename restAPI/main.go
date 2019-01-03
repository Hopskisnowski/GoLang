package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// A Response struct to map the Entire Response
type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

// A struct to map our Pokemon's Species which includes it's name
type PokemonSpecies struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

// A struct to map our berries
type BerryFirmness struct {
	Name string `json:"name"`
	Url  string `json:""`
}

func main() {
	response, err := http.Get("https://pokeapi.co/api/v2/pokemon/")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	// var files string

	// files, _ := ioutil.ReadDir("/")

	// for _, file := range files {
	// 	if file.IsDir() {
	// 		fmt.Print("Dir - ")
	// 	} else if !file.IsDir() {
	// 		fmt.Printf("File - ")
	// 	}
	// 	fmt.Printf("%s - ", file.Name())
	// 	fmt.Printf("%d \n", file.Size())
	// }

	responseData, _ := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	//lastresponseData := responseData

	//diffStr := difference(lastresponseData, responseData)

	// for _, diffVal := range diffStr {
	// 	fmt.Println(diffVal)
	// }

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	// fmt.Println(responseObject.Name)
	// fmt.Println(len(responseObject.Pokemon))

	for i := 0; i < len(responseObject.Pokemon); i++ {
		// if responseObject.Pokemon[i].Species.Name == "gastly" {
		// fmt.Println("url:" + responseObject.Pokemon[i].Species.Url)
		// }
		fmt.Println(responseObject.Pokemon[i].Species.Name + " " + responseObject.Pokemon[i].Species.Url)
	}

}

// func difference(slice1 []byte, slice2 []byte, slice3 []byte) []byte {
// 	diffStr := []byte{}
// 	m := map[byte]int{}

// 	for _, s1Val := range slice1 {
// 		m[s1Val] = 1
// 	}
// 	for _, s2Val := range slice2 {
// 		m[s2Val] = m[s2Val] + 1
// 	}
// 	for _, s3Val := range slice3 {
// 		m[s3Val] = m[s3Val] + 1
// 	}

// 	for mKey, mVal := range m {
// 		if mVal == 1 {
// 			diffStr = append(diffStr, mKey)
// 		}
// 	}

// 	return diffStr
// }
