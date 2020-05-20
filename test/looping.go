package main //must specify a package so that other program can reuse this code by importing 
import "fmt" // fmt stands for format
func main(){ //specify a function with keyword func

	i:=1 //If I declare a variable but not use it, it throws error! This is a shorthand to writing var i =1
	for i<=5{ // Go has only one kind of loop - for loop--No while/do while loops
		fmt.Println("Astha") //Println is a method in fmt package
		i++ //If I don't increment the value of i, it prints my name infinite times
	}
	printhello100times()
	//the following function will print hello 100 times using a different syntax of for loop
}

func printhello100times(){

	for i:=1;i<=100;i++{
		fmt.Println("Hello")
	}
}
