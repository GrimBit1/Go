package main

import "fmt"

func main() {
	println("Welcome to loops in golang ")
	days := []string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
	}
	fmt.Println(days)

	//for loop
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	//for each loop
	// for i := range days {
	// 	fmt.Println(i)
	// }

	//for of loop
	for _, day := range days {
		// index = index
		fmt.Println(day)
	}

	rougeValue := 1

	for rougeValue < 10 { // while loop copy
		// if rougeValue == 5 {
		// 	break
		// } // break

		if rougeValue == 5 {
			rougeValue++
			continue
		}
		fmt.Println(rougeValue)
		rougeValue++

	}

}
