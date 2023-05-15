package main

import "fmt"

func main() {
	fmt.Println("hehehehe")
	var aditya = User{"Aditya"}
	aditya.getName()
}

type User struct {
	Name string
}

func (U User) getName() { // this will make method for User because it is taking first parameter as struct
	fmt.Println(U.Name)

}
