package main


// type creditLimit struct{
// 	accountNumber int64
// 	openingBalance float64
// 	charges float32
// 	credited float64
// 	creditLimits int32
// }

func creditLimitCalculator(openingBalance float64, charges float64, credited float64) float64{
	return openingBalance + charges - credited
}