package main 

import (
	"fmt"
)

func main(){
	myBill := newBill("mario's bill")
	fmt.Println(myBill.format())
	myBill.updateTip(10)
	fmt.Println(myBill.format())
	myBill.addItem("onion soup", 4.50)
	fmt.Println(myBill.format())
}