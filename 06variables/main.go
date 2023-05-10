package main

import "fmt"

func main() {
	var username string = "aditya"
	var age int = 50
	var boolean bool = true
	var number int = 54
	var decimal float32 = 5.01111
	fmt.Println(username, age, boolean, number, decimal)
	//placeholder
	// %T gives type of variable

	fmt.Printf("username is %T \n", username)
	fmt.Println(`username is ${username}`)

	//default values and some aliases
	var othernumber int
	var othername string
	fmt.Println(othernumber)
	fmt.Println(othername)

	// implicit type
	var tempvar = "Hi"
	fmt.Println(tempvar)
	// lexer will automatically add the temp var type as per initialized value
	// but then we cant change it value type

	// no var style
	numberofUser := 5000
	fmt.Println(numberofUser)
}
