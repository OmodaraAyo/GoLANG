package main

import "fmt"

func main(){

	menu := map[string]float64{
		"soup": 4.99,
		"pie": 7.99,
		"salad": 6.99,
		"toffee pudding": 3.55,
	}
	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping through a map
	for k, v := range menu{
		fmt.Println(k, "-", v)
	}

	//int as key type
	phonebook := map[int]string{
		2349094848: "Mario",
		2340193948: "John",
		2341028384: "Mark",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[2340193948])

	for k, v := range phonebook{
		fmt.Println(k, "->", v)
	}

	//update a key value
	phonebook[2341028384] = "Pablo"
	fmt.Println(phonebook)
}