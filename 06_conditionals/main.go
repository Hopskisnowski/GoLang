package main

import (
	"fmt"
)

func main() {
	x := 15
	y := 10

	if x < y {
		fmt.Printf("%d is less than %d \n", x, y)
	} else {
		fmt.Printf("%d is less than %d \n", y, x)
	}

	color := "blue"

	if color == "red" {
		fmt.Printf("color is equal to red\n")
	} else if color == "blue" {
		fmt.Printf("color is equal to blue\n")
	} else {
		fmt.Printf("color is %s\n", color)
	}

	switch color {
	case "red":
		fmt.Printf("Color is red\n")
	case "blue":
		fmt.Printf("Color is blue\n")
	default:
		fmt.Printf("Color is %s\n", color)
	}

}
