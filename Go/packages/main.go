package main

import (
	"fmt"
	"sort"
)

func main(){
	// greeting := "hello go lang"
	// result  := strings.Contains(greeting, "hello") 
	// fmt.Println(strings.Contains(greeting, "hello"))
	// fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	// fmt.Println(strings.ToUpper(greeting))

	// fmt.Println(strings.Index(greeting, "h"))
	// fmt.Println(strings.Split(greeting," "))

	ages := []int{45,20,50,30,75,60,50}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 90)
	fmt.Println(index)
}