package main

import (
	"math/rand"
)

func main() {
	println("In the Switch")
	var randomNum int = rand.Intn(6) + 1
	println(randomNum)
	switch randomNum {
	case 1:
		println("One")
	case 2:
		println("Two")
	case 3:
		println("Three")
	case 4:
		println("Four")
	case 5:
		println("Five")
		// fallthrough
	case 6:
		println("Six")
	default:
		println("Unknown Number")
	}
	// this switch is different from another languages it automatically breaks after 1 case
	// for this we have to fallthrough to also go in another case not break
}
