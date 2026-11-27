package main

import (
	"fmt"
	"math/rand"
) // file block

// package block
var myMap map[int]string

func main() {

	// Shadowing variables example
	name := "pawar"
	fmt.Println(name)
	if name == "pawar" {
		name := "saish"
		fmt.Println(name)
	}
	fmt.Println(name)

	// If statements
	age := 22
	if age > 18 {
		fmt.Println("You are an adult")
	} else if age > 11 {
		fmt.Println("you are a teenager")
	} else {
		fmt.Println("you are a child")
	}

	// If statements with scooping variable to conditon
	if n := rand.Intn(10); n == 0 {
		fmt.Println("Chehh - zero")
	} else if n > 5 {
		fmt.Println("chehh - 5-10")
	} else {
		fmt.Println("CHHHHEHHH")
	}

	i := 100

	/// FOR LOOPS - 4 types
	//
	for i := 0; i < 10; i++ { // we can also skip any part in the condition
	} // standard c-style for
	for i < 100 {
	} // (while loop in go)
	for {
	} /// infinite loop, use break / continue

	// for k, v := range myMap {
	// } // for-range loop
	// for i := range 10 {
	// } // range over an integer -> 0..9

	/// FOR-Range loops

	//for key/index , //value := range <arays/slices/maps/strings/channels/integers/iterator fucntions>  ( if want only key, can ommit value)
	// for k := range ..
	// for _ , v := range ..

	// for strings, range walks runes and not bytes

	// Iterating over a map is not always same order to prevent Hash-DoS attacks

	// SWITCH CASE - GO
	s := "saishpawar11"

	switch size := len(s); size {
	case 1, 2, 3:
		fmt.Println("Too low length")
	case 4:
		fmt.Println("Too medium length")
	case 6, 7, 8, 9:
	default:
		fmt.Println("Long")
	}

}
