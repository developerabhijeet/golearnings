package main

import "fmt"

func main(){
	fmt.Println("If else")

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount>10 {
		result = "watch out"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println("result:", result)

	if 9%2 == 0 {
		fmt.Println("Num is even")
	} else {
		fmt.Println("Num is odd")
	}

	if num:= 3; num<10 { 
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is not less than 10")
	}
	
}