package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hi this is slices")
	var fruitList = []string{"Apple", "Oranges", "Banana", "Peach"}
	fmt.Println(fruitList)
	fmt.Printf("Type of the fruitlist is %T\n", fruitList)
	fmt.Println(fruitList[1:3]) // this is like slicing
	// append()

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777 / this will error because the length is already declared

	// but there is a way

	highScores = append(highScores, 555, 666, 777) // this works because the highscores returns the array and after parameter combine and returns a new array then the array is stored in the same variable
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	// Remove element based on index
	var courses = []string{"Web Development", "JavaScript", "Tailwind CSS", "NodeJs/Express", "VueJS","Solidity"} // this is a slice
	fmt.Println(courses)
	// what if we want to remove an element based on index
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...) // ... is like spread operator in js but it is used after the element
	fmt.Println(courses)

}
