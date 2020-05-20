package main
import "fmt"

var a=9

func demo(){
	a=8
	fmt.Println("In demo: ",a) //function gives preference to its own variable first- function level variable
}	// this is called shadowing
func main(){
	demo()
	fmt.Println("In main: ",a) //package level variable
}
