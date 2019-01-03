package main

import "fmt"

func main() {
	//Define map
	emails := make(map[string]string)

	//Assign

	//Declare map and add kv

	emails["Rasmus"] = "hoppedyt@gmail.com"
	emails["Sanne"] = "sannewb@gmail.com"
	emails["Hans"] = "hans@gmail.com"

	fmt.Println(emails)

	// Delete from mao
	delete(emails, "Hans")

	fmt.Println(emails[Hans])
}
