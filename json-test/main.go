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

// A Response struct to map the Entire Response
type Response struct {
	XCloud    string `json:"X-Cloud-Trace-Context"`
	Upgrade   string `json:"Upgrade-Insecure-Requests"`
	AccLang   string `json:"Accept-Language"`
	Host      string `json:"Host"`
	UserAgent string `json:"User-Agent"`
	Accept    string `json:"Accept"`
	// Time         string `json:"time"`
	// Milliseconds uint   `json:"milliseconds_since_epoch"`
	// Date         string `json:"date"`
}

func main() {
	response, err := http.Get("http://headers.jsontest.com/")
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

	fmt.Println("XCloud: " + responseObject.XCloud)
	fmt.Printf("Upgrade: %s \n", responseObject.Upgrade)
	fmt.Println("AcceptLanguage: " + responseObject.AccLang)
	fmt.Println("Host: " + responseObject.Host)
	fmt.Printf("UserAgent: %s \n", responseObject.UserAgent)
	fmt.Printf("Accept: %s \n ", responseObject.Accept)
	// fmt.Println("Date: " + responseObject.Date)
	// fmt.Printf("Milliseconds: %d \n", responseObject.Milliseconds)
	// fmt.Println("Time: " + responseObject.Time)

	fmt.Println(len(responseData))

	// for i := 0; i < len(responseData); i++ {
	// 	fmt.Println(responseData)
	// }

}
