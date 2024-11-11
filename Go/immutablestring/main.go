package main 

import "fmt"

func mutateString(x *string){
	*x = "Morries" 
}

func main(){

	name := "pablo"

	nameAddress := &name

	mutateString(nameAddress)

	fmt.Println(name)
}