package main

import(
	"fmt"
)

// func validateInput(input float64)float64{
// 	for{
// 		value, err := fmt.Scan(&input)
// 		if err != nil{
// 			fmt.Println("Please provide a number")
// 			continue
// 		}else{
// 			return float64(value)
// 		}
// 	}
// }

func main(){
	var accountNumber, openingBalance, charges, credited, creditLimit float64

	fmt.Println("Enter account number: ")
	fmt.Scan(&accountNumber)

	fmt.Println("Enter credit limit: ")
	fmt.Scan(&creditLimit)	

	fmt.Println("Enter opening balance: ")
	fmt.Scan(&openingBalance)

	fmt.Println("Enter total Items charges: ")
	fmt.Scan(&charges)

	fmt.Println("Enter total of all credits: ")
	fmt.Scan(&credited)

	newBalance := creditLimitCalculator(openingBalance, charges, credited)
	if newBalance > creditLimit{
		fmt.Println("Credit card limit exceeded")
	}else{
		fmt.Printf("New Balance: $%0.2f",newBalance)
	}
}