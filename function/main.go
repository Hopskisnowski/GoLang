package main

import "fmt"

func PrintName(First string, Last string) (string, string) {
	return First, Last
}

func main() {
	//errorMessage := "No error"
	fmt.Println("Hello World")
	result1, result2 := PrintName("Hans", "Hansen")
	//if err != nil {
	fmt.Printf("%s %s \n", result1, result2)
	//}
	//fmt.Println(result)
}
