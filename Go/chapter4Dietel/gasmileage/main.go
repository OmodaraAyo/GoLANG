package main

import (
	"fmt"
	
)

func main(){
	allTripsMileage := []float64{}
	var milesDriven, gasUsed float32

	var opt int
	for{
		fmt.Println("Select an option\n1: calculate vehicle mileage\n2: Exit")
		_, err:= fmt.Scan(&opt)

		if err != nil{
			fmt.Print("Invalid input, please enter a valid number\n\n")
			continue
		}

		switch opt{
			case 1:
				result := runApp(milesDriven, gasUsed)
				fmt.Printf("Your vehicle mileage = %0.2f\n\n", result)
				allTripsMileage = append(allTripsMileage, result)
			case 2: 
				fmt.Printf("Combined mileage for all trip is: %0.2f\n", totalMileage(allTripsMileage))
				fmt.Println("App closed")
				return
			default:
				fmt.Printf("Please Select a valid option\n\n")
		}
	}
}

// var milesDriven int32
// 		fmt.Print("Enter miles driven: ")
// 		fmt.Scan(milesDriven)

// 		var gallon int32
// 		fmt.Print("Enter gallon used: ")
// 		fmt.Scan(gallon)
		
// var num int
// 	fmt.Scan(&num)
// 	fmt.Println("you number is =",num)


// scanner := bufio.NewScanner(os.Stdin)
// 	fmt.Println("Enter text or -1 to exit")

// 	for {
// 		fmt.Print("->")
// 		scanner.Scan()
// 		input := strings.TrimSpace(scanner.Text())
		
// 		if input == "-1"{
// 			fmt.Println("exit")
// 			break
// 		}
// 	}