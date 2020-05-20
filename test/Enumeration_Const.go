package main

import "fmt"

const (
	isAdmin = 1 << iota // multiply by 2**0=1
	isHeadquarters // using enumeration, this value will be 2**1

	canseefinancials//this value will be 2**2
)
func main(){

	var roles byte = isAdmin | isHeadquarters | canseefinancials
	fmt.Printf("%b\n", roles) //prints 111= basically the OR result of 00000001=admin, 00000010=headquarter, 00000100=canseefinancials

	fmt.Printf("Is Admin? %v\n", isAdmin & roles== isAdmin)


}