package main

import(
	"fmt"
	"time"
)

func main() {
	fmt.Println("When is Friday?")
	today := time.Now().Weekday()
	switch time.Friday {
	case today + 0:
		fmt.Println("It is today!") //no need of break statement
	case today + 1:
		fmt.Println("In 1 day")
	case today + 2:
		fmt.Println("In 2 days")
	default: 
		fmt.Println("too far away!")
		
	}
}