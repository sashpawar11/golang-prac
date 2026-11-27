package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Sum of 2+2 = ", addNum(2, 2))

	fmt.Println("total of first 10 natural numbers: ", getTotalSum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	numSlice := []int{100, 200, 300, 400, 500}
	fmt.Println("passing slice as a variadic input to getTotal():  ", getTotalSum(numSlice...))

	fmt.Println("Func with multiple return values: ")
	result, remainder, err := divAndRemainder(2, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result of divide and remainder\n Result : ", result, " Remainder: ", remainder)
	}

}

// func <funcnam>(<params : type>) <returntype> {}
func addNum(a int, b int) int {
	return a + b
}

// Go does nothave named and optional parameters -> can be simulated by passing struct as the functon param

// Variadic parameters :
// Rules : should be the last paarameter in a function, and employ ... dots/
// The parameter is treated as a slice, so any slice functions or operations are supported on it
func getTotalSum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

/// Functions with multiple return values

func divAndRemainder(num int, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("Err: Cannot divide by zero")
	}
	return num / denom, num % denom, nil
}
