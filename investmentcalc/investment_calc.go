package main

import (
	"fmt"
	"math"
)

// investment calculator using user inputs
// camelCase for variables
func main() {
	fmt.Print("Investment Calculator!")

	//variables
	var investmentAmount = 1000
	var roi = 5.5
	var years = 10
	var futureAmount = float64(investmentAmount) * math.Pow(1+roi/100, float64(years))

	fmt.Print("Final Amount: ", futureAmount)

}
