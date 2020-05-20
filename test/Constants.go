package main

import "fmt"

const a int8=10

func constshadowing(){
	const a int32 = 70//even though a is a package level constant, we can change its value using the concept of shadowing
	fmt.Printf("\n%v %T",a,a)
}
func addconstandvar(){
	const d int=2
	var b int=8
	fmt.Printf("\n%v %T", b+d,b+d)//we can add var and const and result will be var
}
func adddiffinttypes(){
//	const d int=2
//	var b int32=8
	//fmt.Printf("\n%v %T", b+d,b+d)//this will give error. cant add different types. no concept of implicit type conversion. always explicit
}

func main(){

	const a int = 14
	const b string = "astha"
	const c float32 = 3.14
	const d bool = true
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	constshadowing()
	addconstandvar()
	adddiffinttypes()
}