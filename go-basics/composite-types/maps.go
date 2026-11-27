package main

import (
	"fmt"
)

func main() {

	// MAPS data structrue - helps store key-value pairs
	// Zero Value = nil
	// map[key]value
	// Internally, go uses hash maps or hash tables to implement maps

	//Declaration

	var nilMap map[int]string // key int type, values string type

	fmt.Println(nilMap)
	fmt.Printf("type of nilMap, %T\n", nilMap)
	//nilMap[1] = "saish" // panic, writes not allowed in nil type map declaration

	fmt.Println("Does nilMap == nil?:", nilMap == nil)

	// Empty map declartion  -- /writes allowed, reads show elements
	map2 := map[int]string{}
	map2[1] = "saish"
	map2[2] = "pawar"

	fmt.Println(map2, len(map2))

	// declaring non-empty map
	groups := map[string][]string{
		"groupA": {"India", "Australia"},
		"groupB": {"SA", "New Zealand"},
		"groupC": {"England", "Srilanka"},
	}

	fmt.Println(groups)
	// compare using maps.Equal or maps.EqualFunc
	//

	///****
	//// Checking if a key exists using ""comma ok idom"""

	colors := map[string]string{}
	colors["red"] = "#ff0000"
	colors["green"] = "#00ff00"
	colors["black"] = "#000000"
	colors["white"] = "#ffffff"

	v, ok := colors["red"] /// commma ok idiom
	// if i dont need value, i can do -->  _ , ok := colors["red"]
	fmt.Println("Color red exists in map? : ", ok)

	// accessing it returns zero value for the non-existing key
	// so colors["orange] would return "" (zero value for strings)
	if ok {
		fmt.Println("Color value: ", v)
	}

	//Deleting in a map
	//
	delete(colors, "white")
	fmt.Println(colors)

	// Emptying a map
	//
	//clear(colors)
	fmt.Println(colors)

	// Maps are compares using maps.Equal function
	//fmt.Println(maps.Equal(map2, colors))
	//
	//

	// Go does not have SET, but you can use maps to make a SET
	/*
		* A set is a data type that ensures
					there is at most one of a value, but doesn’t guarantee that the values are in any
					particular order.
	*/

	//
	intSet := map[int]bool{}
	// quick way to fill the map keys
	intSlces := []int{5, 200, 245, 1, 23, 543, 674, 342, 21, 2452}

	for _, v := range intSlces {
		intSet[v] = true
	}

	// we now have a map whose keys are ints and their values are bool = true, so any key not present would return the zero value for bool type i.e. false!
	//
	fmt.Println("205 present in set? : ", intSet[205])
	fmt.Println("5 present in set? : ", intSet[5])

}
