package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	// no inheritance in golang; No super or parent
	abhijeet := User{"Abhijeet", "abhi@tech.com", true, 24}

	fmt.Println("USer:", abhijeet)
	fmt.Printf("details %+v\n", abhijeet)
	fmt.Printf("Name is: %v.", abhijeet.Name)

	abhijeet.GetStatus()
	abhijeet.NewMail()

}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}

func (u User) GetStatus(){
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail(){
	u.Email = "abh@hash.co"
	fmt.Println("Email of this user:", u.Email)
}