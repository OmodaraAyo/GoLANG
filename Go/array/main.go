package main

import "fmt"

func main(){

	var ages [3]int = [3]int{20, 25, 30}
	fmt.Println(ages)
	
	var age = [3]int{20, 25, 30}
	fmt.Println(age, len(age))

	name := [4]string{"ayo", "dele", "mark", "pablo"}
	fmt.Println(name, len(name))

	//slice (use arrays under the hood) can append using slice
	var scores = []int{2,4,5}
	fmt.Print(scores, len(scores), "\n\n")
	scores[2] = 25
	fmt.Print(scores, len(scores), "\n\n")
	scores = append(scores, 85)
	fmt.Print(scores, len(scores), "\n\n")

	//slice range
	rangeOne := name[1:3]
	fmt.Print(rangeOne,"\n\n")

	rangeTwo := name[2:]
	fmt.Print(rangeTwo, "\n\n")

	rangeThree := name[:3]
	fmt.Print(rangeThree, "\n\n")

	rangeOne = append(rangeOne, "koppa")
	fmt.Print(rangeOne, "\n\n")
}