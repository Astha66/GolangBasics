package main

import "fmt"

func main(){
 s:= "this is a string"
 s1:= " also a string"
 fmt.Println(s+s1) //CONCATENATION

 //s[2]='i' //not allowed= string is immutable
//s[2]= can't directly access a string like this. Rather we use the string() function
fmt.Printf("%v %T", string(s[2]), s[2])

b := []byte(s)
fmt.Printf("%v %T",b,b)
}
//conversion from string to byte= []byte(s)
//byte to string= string(b)