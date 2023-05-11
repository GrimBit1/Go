package main

import "fmt"

func main() {
	fmt.Println("Hi this is for pointers")
	var b *int // Default value of pointer is nil
	fmt.Println(b)
	// & means reference
	// * means the value in the value in the pointer
	var a = 12
	b = &a
	fmt.Println(b)
	*b = 11
	fmt.Println(a)
}
