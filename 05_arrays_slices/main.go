package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// // Arrays
	// var fruitArr [2]string

	// // Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Declade and assign
	// fruitArr := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[0])
	// fmt.Println(fruitArr[1] + " happy " + fruitArr[0])
	fruitSlice := []string{"Apple", "Orange", "Grape"}
	fmt.Println(fruitSlice)

	response, err := http.Get("https://httpbin.org/ip")
	if err != nil {
		fmt.Println("The HTTP request failed \n")
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
	jsonValue, _ := json.Marshal(jsonData)
	response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
	// fmt.Println(fruitSlice[0])
	// fmt.Println(fruitSlice[1])
	// fmt.Println(fruitSlice[2])
}
