package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader)(string, error){
	fmt.Print((prompt))
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}
func createBill()bill{
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("create a new bill name: ")

	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)
	//to one line

	name, _ := getInput("create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill -", b.name)

	return b
} 

func promptFunction(b bill){
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option \na - add items\ns - save bill\nt - add tip\n ", reader)
	
	switch opt{
		case "a":
			name, _ := getInput("Item name: ", reader)
			price, _ := getInput("Item price: ", reader)
			p, err := strconv.ParseFloat(price, 64)
			if err != nil{
				fmt.Println("price must be a number")
				promptFunction(b)
			}
			b.addItem(name, p)
			fmt.Println("Item added",name, p)
			promptFunction(b)
		case "t":
			tip, _ := getInput("Enter tip amount ($): ", reader)
			t, err := strconv.ParseFloat(tip, 64)
			if err != nil{
				fmt.Println("tip must be a number")
				promptFunction(b)
			}
			b.updateTip(t)
			fmt.Println("Tip added - ",tip)
			promptFunction(b)
		case "s":
			b.save()
			fmt.Println("file saved - ", b.name)
		default:
			fmt.Println("not a valid option")
			promptFunction(b)
	}
}

func main(){
	myBill := createBill()
	promptFunction(myBill)

}