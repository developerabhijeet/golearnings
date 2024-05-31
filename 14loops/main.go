package main

import "fmt"

func main() {
	fmt.Println("Welcome to golang")
	days := []string{"sunday", "tuesday", "wednesday", "friday", "saturday"}

	fmt.Println("days:", days)

	for d := 0; d < len(days); d++ {
		fmt.Println("d:", days[d])
	}

	for i:= range days{
		fmt.Println("days: ",i, days[i] )
	}

	for index, day := range days {
		fmt.Printf("Index is %v and value is %v\n", index, day)
	}

	rougueValue := 1
	for rougueValue < 10 {
		if rougueValue == 4 {
			rougueValue++
			continue
		}else if rougueValue == 6 {
			rougueValue++
			break
		}else if rougueValue==5{
			goto lco
		}
		fmt.Println("Value is:", rougueValue)
		rougueValue++
	}

	lco:
		fmt.Println("I am called by goto statement")

}
