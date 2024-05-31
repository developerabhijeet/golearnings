package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the Rating of our Pizza: ")

	// comma ok || err ok
	// here we call is this as above because go doesn't wants to handle error as it should be
	// It just want to treat like true or false and show accordingly
	// here input is your try and err is catch

	input,err := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
		fmt.Println("Fails, ", err)

}
