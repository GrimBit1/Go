package main

import "fmt"

func main() {
	println("In the structs")
	// no inheritance in golang ; No super or parent
	aditya := User{"Aditya", "adityanandwana01@gmail.com", true, 18}
	fmt.Println(aditya)
	fmt.Printf("%+v\n",aditya)
	fmt.Printf("Aditya details are: %T\n", aditya)

	//to access the child
	fmt.Println(aditya.Name)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
