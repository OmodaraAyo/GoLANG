package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64

}

func newBill(name string) bill {

	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b

}

// format bill
//receiver function
func (b *bill) format() string {
	fs := "Bill breakdown \n"

	var total float64 = 0

	//list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ... $%v\n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%-25v ...$%v\n", "tip: ", b.tip)
	
	//total
	// fs += fmt.Sprintf("total: ...$%0.2f", total)
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total: ", total+b.tip)
	return  fs
}

// how to derefence a struct
func (b *bill) updateTip(tip float64){
	b.tip = tip
} 

//add an item to the bill
func (b *bill) addItem(name  string, price float64){
	b.items[name] = price
}

func (b *bill) save(){
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil{
		panic(err)
	}
	fmt.Println("bill was saved to file")
}