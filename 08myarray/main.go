package main

import "fmt"

func main() {
	fmt.Println("Welcome to array")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "banana"
	fruitList[3] = "tomato"

	fmt.Println("FruitList: ", fruitList)
	fmt.Println("length of FruitList: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushrooom"}
	fmt.Println("List of veg is:", vegList, len(vegList))
}