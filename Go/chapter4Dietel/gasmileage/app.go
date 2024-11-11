package main

import (
	"fmt"
)
// func totalMileage(value float64){
// 	allTripsMileage := []float64{}
// 	allTripsMileage = append(allTripsMileage, float64(value))
// 	fmt.Println(allTripsMileage)
// }

func runApp(milesDriven float32, gasUsed float32) float64{
		fmt.Println("Enter miles driven: ")
		fmt.Scan(&milesDriven)

		fmt.Println("Enter gas used: ")
		fmt.Scan(&gasUsed)

		gasMileage := milesDriven / gasUsed
		return float64(gasMileage)
}

func totalMileage(value []float64) float64{ 
	var total float64
	for i := 0; i < len(value); i++{
		total += value[i]
	}
	return total
}
