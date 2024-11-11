package main

import (
	"fmt"
)
func main(){
	var totalSale, count, value int64

	var numberOfItemSold int
	fmt.Println("Enter number of item sold: ")
	fmt.Scan(&numberOfItemSold)

	if validateInput(numberOfItemSold) != numberOfItemSold{
		return
	}

	for i := 0; i < numberOfItemSold; i++{
		count++
		fmt.Printf("Enter Item %v value: $", count)
		fmt.Scan(&value);
		totalSale += value;
	}
	fmt.Printf("Your salary is: $%0.2f", salesCommissionCalculator(float64(totalSale)))
}