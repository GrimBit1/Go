package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "mango"
	fruitList[3] = "banana"
	fmt.Println(fruitList)
	// even there is 3 elements in the array the len is 4 because we have already declared its length
	fmt.Println(len(fruitList))
	var vegList = [3]string{"Hi", "imnot a fruit", "hehehe"}
	fmt.Println(vegList)
}
