package main
import(
	"fmt"
	"math"
)// this is how we import multiple packages in go
func main(){
	var num float64= 9
	var num1 float64 = 12
	var result = math.Sqrt(num)
	var result1 = math.Sqrt(num1)
	var intResult = math.Round(result1)
	var floorResult = math.Floor(result1)
	var ceilResult = math.Ceil(result1)
	fmt.Println(result)
	fmt.Printf("The output is %g \nThank you\n",result1)//we can use both Printf and Println. If i want only 2 places after decimal, %0.2g
	fmt.Println(intResult)
	fmt.Println(floorResult)
	fmt.Println(ceilResult)
}
