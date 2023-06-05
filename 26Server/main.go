package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Hi")
	// PerformGetRequest()
	// PerformPostJSONRequest()
	PerformPostRequest()
}

func PerformGetRequest() {
	const url = "http://localhost:3000/"
	res, err := http.Get(url)
	checkError(err)
	resdata, err := io.ReadAll(res.Body)
	checkError(err)
	fmt.Println(string(resdata))
	fmt.Println("Performing Get Request")
	defer res.Body.Close()
}

func PerformPostJSONRequest() {
	const url = "http://localhost:3000/post"
	const json = `{"name":"John"}`
	fmt.Println(json)
	var body = strings.NewReader(json)
	res, err := http.Post(url, "application/json", body)
	checkError(err)
	defer res.Body.Close()
	data, _ := io.ReadAll(res.Body)
	fmt.Println(string(data))
	// fmt.Println("Performing Post Request")
	// fmt.Println(url)
}
func PerformPostRequest() {
	const myurl = "http://localhost:3000/postform"
	data := url.Values{}
	data.Add("name", "Aditya")
	res, err := http.PostForm(myurl, data)
	checkError(err)
	defer res.Body.Close()
	mydata, _ := io.ReadAll(res.Body)
	fmt.Println(string(mydata))

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
