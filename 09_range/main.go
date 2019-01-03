package main

import "fmt"

func main() {
	ids := []int{33, 76, 48, 39, 12, 9, 2}

	//Loop through ids
	// for i, id := range ids {
	// 	fmt.Printf("%d - ID : %d \n", i, id)
	// }

	sum := 0
	for _, id := range ids {
		fmt.Printf("%d \n", id)
		sum = id + sum
		fmt.Printf("Sum - %d \n", sum)
	}

	//Range with map

	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s \n", k, v)
	}
}
