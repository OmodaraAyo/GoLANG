package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string){
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var intials [] string

	for _, v := range names{
		intials = append(intials, v[:1])
	}
	if len(intials) > 1{
		return intials[0], intials[1]
	}
	return intials[0], "_"
}
func main(){
	fn, sn := getInitials("tifa lockhart")
	fmt.Println(fn, sn)

	fn2, sn2 := getInitials("Garath")
	fmt.Println(fn2, sn2)
}