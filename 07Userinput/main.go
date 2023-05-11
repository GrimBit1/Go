package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Hello"
	fmt.Println(welcome)
	//comma ok /comma error syntax
	var reader = bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	input, err := reader.ReadString('\n') // this is like a try catch in js but because golang treats err also as value it doesn't effect runtime
	fmt.Println(input, err)
	fmt.Printf("Type of the input is %T \n", input)
	fmt.Print(reader.ReadString)
}
