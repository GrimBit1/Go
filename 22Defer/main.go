package main

import "fmt"

func main() {
	defer hi()
	hehe() // the function calls are done syncronously but if i use defer then it will run before returning
}

func hi() {
	fmt.Println("Hi")
}

func hehe() {
	fmt.Println("hehehe")
}
