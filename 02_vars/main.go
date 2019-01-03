package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using the var keyword

	name := "Rasmus"

	var size float32 = 2.3

	username, email := "rsb", "@bec.dk"

	age := 44
	isCool := true
	isCool = false
	size = 1.3

	fmt.Println(name, age, isCool, size)
	fmt.Printf("The variable name holds a %T type and the variable age holds a %T type and finally my shoesize is of type %T\n", name, age, size)
	fmt.Printf("My name is %s and i'm %d years old and %T \n", name, age, isCool)
	fmt.Printf("username is %s and emain is %s \n", username, email)
}
