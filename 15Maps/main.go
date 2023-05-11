package main

import "fmt"

func main() {
	fmt.Println("In the maps")
	languages := make(map[string]string, 4)
	fmt.Println(languages)
	languages["js"] = "JavaScript"
	languages["rb"] = "Ruby"
	languages["sol"] = "Solidity"

	//delete keyword to delete a key value pair
	delete(languages, "rb")
	fmt.Println(languages)

	//loops are interesting in golang

	for k, v := range languages {
		fmt.Printf(k, v, "\n")
		fmt.Printf("Hi the key %v for value %v \n", k, v)
	}

}
