package main

import "fmt"

func main() {

	// ARRAYS IN GO  -- not recommended or used explictly usually in go due to their type restrictions and conversions

	// Style1 : size first and type of elements in the arr
	var x [3]int // values initialized ( {0,0,0})

	///// to initialize values to the arr
	var x = [3]int{1, 2, 3}

	// to create sparse array ( most elements are zero with some non-zero elements)
	var sparseArr = [12]int{1, 5, 4: 12, 9: 100, 15}
	// {1, 5, 0 , 0 , 0, 12, 0 ,0, 0, 0, 100, 15}

	// * Style 2
	var y = [...]int{5, 24, 51, 66, 21, 4} /// [...] can be used to not specifiy numb3e of elemenets

	// 2D Arrays
	var twodArr = [2][3]int
	// a b c
	// d e f

	// built-in length function for arrays
	fmt.Println(len(sparseArr))
}
