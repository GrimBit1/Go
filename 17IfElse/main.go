package main

import "fmt"

func main() {
	println("hehehehehe")
	loginCount := 10
	var result string
	if loginCount < 10 {
		result = "Regular User"

	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10"
	}
	fmt.Println(result)
}
