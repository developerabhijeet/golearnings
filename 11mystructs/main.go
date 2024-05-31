package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	// no inheritance in golang; No super or parent
	abhijeet := User{"Abhijeet", "abhi@tech.com", true, 24}

	fmt.Println("USer:", abhijeet)
	fmt.Printf("details %+v\n", abhijeet)
	fmt.Printf("Name is: %v.", abhijeet.Name)

}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}
