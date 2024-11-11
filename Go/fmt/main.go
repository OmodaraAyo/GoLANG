package main

import "fmt"

func main(){
	age := 25
	name := "gucchie"
	fmt.Print("hello, \n")
	fmt.Print("world! ")

	fmt.Println("My name is", name, "and i'm", age, "years old.")

	//to create a string with variable embedded inside(formatted String) %_ = format specifier
	fmt.Printf("my age is %v and my name is %v\n", age, name) 

	//puts quotes around the string.
	fmt.Printf("my age is %q and my name is %q\n", age, name) 

	// of what data type
	fmt.Printf("age is of type %T\n", age)

	// 
	fmt.Printf("your score %0.2f points!\n", 225.55)

	//sprintF (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v\n", age, name)
	fmt.Println("the saved string is: ",str)
}