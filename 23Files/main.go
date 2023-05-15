package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("hehehe")
	content := "heheehehhehehehe"
	file, _ := os.Create("temp.txt")
	fmt.Println(file)
	sucess, err := file.WriteString(content)
	if err != nil {
		fmt.Println("Got an error")
		panic(err)
	}
	fmt.Println(sucess)
	data, _ := ioutil.ReadFile("./temp.txt")
	fmt.Println(string(data))

}
