package main

import "fmt"

func main() {
	fmt.Println("Hi")
	result := add(1, 2)
	fmt.Println(result)
	fmt.Println(proAdder(1, 2, 3, 4, 56, 4, 6, 5, 78, 5, 7))
}
func add(valueOne int, valueTwo int) int { // we have to which type of parameter we are expecting and which type of value we are returning
	return valueOne + valueTwo
}
func proAdder(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}
