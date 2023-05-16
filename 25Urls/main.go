package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://jsonplaceholder.typicode.com/photos/20?he=he"

func main() {
	fmt.Println("Urls")
	fmt.Println(myurl)
	temp, err := url.Parse(myurl)
	checkError(err)
	fmt.Printf("%T\n", temp)
	fmt.Println(temp.Scheme)   //Gives protocol of the url
	fmt.Println(temp.Port())   //Gives port of the url /if blank means 80 or 443
	fmt.Println(temp.Host)     //Gives host of the url
	fmt.Println(temp.Path)     //Gives path of the url / after the hostname path
	fmt.Println(temp.RawQuery) //Gives query of the url
	var qparams = temp.Query() // makes a map of all the queries
	fmt.Println(qparams)
}
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
