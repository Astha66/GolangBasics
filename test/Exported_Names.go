package main

import "fmt"


func main(){ //is only accessible within this package

	fmt.Println("astha")
	Demo()
}

func Demo(){				// can be accessed through different packages- name starts with CAPS just as in case of Println
	fmt.Println("Sharma")
}