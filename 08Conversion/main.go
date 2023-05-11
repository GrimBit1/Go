package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// there is no exception error in the golang
func main() {
	fmt.Println("hi welcome to our restaurant app")
	fmt.Println("Please rate our app")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Rate")
	input, _ := reader.ReadString('\n')
	number, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	if err != nil { // if else statement
		fmt.Println(err)
	} else {
		fmt.Println(1 + number)
	}
}
