package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")


	var fruitList = []string{"apple", "tomato", "orange", "peach"}

	fmt.Printf("type of fruitlist is %T\n", fruitList) 
	fruitList = append(fruitList, "Mango", "banana")

	fmt.Println("Fruitlist:", fruitList)

	fruitList = append(fruitList[1:])


	fmt.Println("sliced:", fruitList)
	fruitList = append(fruitList[1:3])

	fmt.Println("sliced ranged:", fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println("sliced last:", fruitList)

	highScores := make([]int, 4)

	highScores[0] = 123
	highScores[1] = 125
	highScores[2] = 126
	highScores[3] = 873

	fmt.Println("show highscores", highScores)

	highScores = append(highScores, 111, 555)
	fmt.Println("highscore new:", highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}

	fmt.Println("courses:", courses)

	// we have removed swift the index 2
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("courses new", courses)


}