package main

import "fmt"

// constant
const PublicToken = "asdfgh"
func main() {
	var username string = "Abhijeet"
	fmt.Println("username: ", username)
	fmt.Printf("variable type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println("Login status:", isLoggedIn)
	fmt.Printf("variable type: %T \n", isLoggedIn)

	// default values and aliases
	var anotherVariable int
	fmt.Println("anotherVariable", anotherVariable)

	// implicit type 
	var my_name = "Abhijeet Tiwari"
	fmt.Println("Name:", my_name)

	// No var style
	number_of_users := 300000
	fmt.Println("Number of users", number_of_users)

	fmt.Println("Constant", PublicToken)
}
