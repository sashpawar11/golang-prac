package main

import "fmt"

func main() {
	fmt.Printf("Hello. %s!\n", "world")
	fmt.Println("Greetings and \n Salutations")

	var test = "sash"
	test = "new"
	fmt.Println(test)

	// rune used when referring to a characte ( alias for int32 type)
	// var myFirstInitial rune = 'S'
	// var myLastInitial rune = 'H'

	// go doesnt support automatic type promotion, explicit type conversion requied

	var x int = 2
	var y float64 = 365.13
	var sum1 float64 = float64(x) + y
	var sum2 int = x + int(y)
	fmt.Println(sum1, sum2)

	var testByte byte = 2000 // overflow ( size limitations on type (byte = uint8 (0-255)))

	//Variable declartion var v/s :=

	// := used for dynamic type variables ( automatically infered types)
	myVariable := true

	var z int /// infers zero value (0)

	var a, b int = 10, 20 // declaring multiple variables at once

	c, d := 10, 30 // declaring multiple variables at once using :=

	//***
	// USING := is illeagal outside functions, e.g. package level etc.
	//***

	//const types

	const myName string = "pawar"
	const pi = 3.14 // prefer keeping constants untyped as it provides flexibility in assignment

	//declaration list style(easily declare multiple variables at once)
	const (
		idKey   = "id"
		nameKey = "name"
	)

	var m = pi //legal assignments

	var (
		variable1 int = 1
		vaiable2      = 3
	)

	// Idiomatic go uses camelCase for varable naming, avoid snake case (e.g. index_counter)
}
