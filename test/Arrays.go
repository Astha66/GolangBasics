package main

import "fmt"

func twoDarrayexample(){
	var identitymatrix [3][3]int
	identitymatrix[0] = [3]int{1,0,0}
	identitymatrix[1] = [3]int{0,1,0}
	identitymatrix[2] = [3]int{0,0,1}

	fmt.Printf("%v\n",identitymatrix)
}
func slicing(){
var a = [...]int{1,2,4}
b :=a
b[1] = 3
fmt.Println(a)//prints 1,2,4
fmt.Println(b)//prints 1,3,4

//however if we dont mention the ... inside [], it is considered as a slice and b will not be a separate array

var p = []int{1,2,4}
q :=p
q[1] = 3
fmt.Println(p)//prints 1,3,4
fmt.Println(q)//prints 1,3,4

var c = []int{1,2,3,4,5,6,7,8,9,10}
fmt.Println(c[:])
fmt.Println(c[3:6])//same as python
fmt.Println(c[:6])

}

func main(){

	var students [3]string
	var marks = [...]int{91,92,93} //another type of declaration
	students[0]= "Astha"
	students[1]= "Pragati"
	students[2]= "sanjana"

	fmt.Printf("%v\n", students)
	fmt.Printf("%v\n", marks)

	//using built in length function
	fmt.Printf("No of students: %v\n",len(students))//prints 3
	//using capacity
	fmt.Printf("capacity: %v\n",cap(students))//prints 3
	twoDarrayexample()
	slicing()
}