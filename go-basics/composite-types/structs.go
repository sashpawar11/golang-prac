package main

import "fmt"

// // STRUCTS in go
type animal struct {
	name  string
	legs  int
	color string
}

func main() {

	// Declaing struct variables
	// 2 Styles for struct fields
	//
	var zoo = animal{
		"tiger",
		4,
		"red",
	}

	// named fields
	//
	var park = animal{
		name:  "cat",
		legs:  4,
		color: "white",
	}

	// using :=

	jungle := animal{
		"lion",
		4,
		"orange",
	}

	fmt.Println(zoo, park)
	fmt.Println(jungle)
	fmt.Println(jungle.color)

	////// ***
	// ANONYMOUS STRUCTS
	//
	var employee struct {
		name string
		age  int
	}

	employee.name = "Sash"
	employee.age = 20

	fmt.Println(employee)

	vehicles := struct {
		model string
		brand string
	}{
		model: "r8 spider",
		brand: "audi",
	}

	fmt.Println(vehicles)
}
