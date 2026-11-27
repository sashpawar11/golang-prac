package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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

	fmt.Println("Func with multiple named return values: ")
	x, y, z := divAndRemainder(2, 2)
	if z != nil {
		fmt.Println(z)
	} else {
		fmt.Println("Result of divide and remainder\n Result : ", x, " Remainder: ", y)
	}

	// Calculator using functions as values concept

	//Slice results map:
	results := map[string]int{}

	// supply function signature as value type for map (function are values)
	opMap := map[string]func(int, int) int{
		"+": add, /// add fucn as valu
		"-": sub, // sub function as value
		"*": mul, // mul func as value
		"/": div, // div func as value
	}

	expressions := [][]string{
		{"2", "+", "4"},
		{"10", "/", "100"},
		{"1", "*", "0"},
		{"20000", "+", "4"},
		{"4"},
		{"20", "%", "4"},
	}

	for _, exp := range expressions {
		if len(exp) != 3 {
			fmt.Println("Invalid expression length - ", len(exp))
			continue
		}
		num1, err := strconv.Atoi(exp[0]) // converting string to int using strconv.Atoi()
		if err != nil {
			fmt.Println(err)
			continue
		}
		opSymbol := exp[1]
		opFunc, ok := opMap[opSymbol]
		if !ok {
			fmt.Println("Invalid operator supplied in expression: ", exp)
			continue
		}

		num2, err := strconv.Atoi(exp[2])
		if err != nil {
			fmt.Println(err)
			continue
		}

		results[strings.Join(exp, " ")] = opFunc(num1, num2)
	}

	fmt.Println("\n****Result of Calculator expressions :\n ")
	for k, v := range results {
		fmt.Println("Expression: ", k, " Result: ", v, "\n")
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
	return num / denom, num % denom, nil // for no errors use nil value
}

// / Functions with named return values
func divAndRemainderWNRP(num int, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		err = errors.New("Err: Cannot divide by zero")
		return result, remainder, err
	}
	result, remainder = num/denom, num%denom
	return result, remainder, nil
}

// Function are values in go - example buildign a calculator
//

// Operation functions
func add(a int, b int) int { return a + b }
func sub(a int, b int) int { return a - b }
func mul(a int, b int) int { return a * b }
func div(a int, b int) int { return a / b }
