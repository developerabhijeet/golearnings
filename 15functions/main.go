package main 

import "fmt"

func main(){
	fmt.Println("I am function")
	greeter()

	result := adder(3,5)
	fmt.Println("result", result)

	proResult, _ := proAdder(2,3,4,5,6)
	fmt.Println("pro adder", proResult)
}

func greeter(){
	fmt.Println("Hello Go")
}
func adder(valOne int, valTwo int) int{
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "Hi I am string"
}