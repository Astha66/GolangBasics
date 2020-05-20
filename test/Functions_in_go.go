package main
import "fmt"

func add(x int,y int) int{
	return x+y
}
func calc (x int,y int) (int,int){ //we can return multiple values unlike c/java
	out1:= x+y
	out2:= x-y
  return out1,out2
}
/*
another format for above function
func calc(x,y int) (out1,out2 int){
	out1 = x+y
	out2= x-y
	return
}
*/

func main(){
	num1:=2
	num2:=3
	result:= add(num1,num2)
	fmt.Println(result)

	resultadd,resultsub := calc(num1,num2)
	fmt.Println(resultadd,resultsub)
}