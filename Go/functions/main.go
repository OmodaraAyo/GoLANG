package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string){
	fmt.Printf("Good morning %v \n", n)
}
func sayBye(n string){
	fmt.Printf("Good Bye %v \n", n)
}
func cycleNames(n []string, f func(string)){
	for _, v := range n{
		f(v)
	}
}

func cycleArea(r float64) float64{
	return math.Pi * r * r
}

func main(){
	// sayGreeting("mario")
	// sayGreeting("Pablo")
	// sayBye("mario")
	// sayBye("Pablo")



	// cycleNames([]string{"Clyde", "pablo", "baron"}, sayGreeting)
	// cycleNames([]string{"Clyde", "pablo", "baron"}, sayBye)


	a1 := cycleArea(10.5)
	a2 := cycleArea(15)
	fmt.Println(a1, a2)
	fmt.Printf("Circle 1 = %0.3f\nCircle 2 = %0.3f", a1, a2)
}