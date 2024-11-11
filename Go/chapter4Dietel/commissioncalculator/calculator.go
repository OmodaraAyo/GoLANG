package main

import "fmt"

func validateInput(number int) int {
	if number <= 0 {
		fmt.Println("Invalid number")
	} else {
		return int(number)
	}
	return 0
}

func salesCommissionCalculator(totalAmount float64) float64 {
	result := 200 + (0.09 * totalAmount)
	return result
}