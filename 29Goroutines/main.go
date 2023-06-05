package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // Creating a wait group

func main() {
	fmt.Println("Hi")
	var urllist = []string{"https://jsonplaceholder.typicode.com/photos/20", "http://google.com", "https://techblogs.codes", "http://fb.com"}
	for _, v := range urllist {
		go getStatusCode(v)
		wg.Add(1) // Add for wait
	}
	defer wg.Wait() // this means not to close main until thread returns
}
func getStatusCode(url string) {
	defer wg.Done() // this means function is done
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Printf("%s got from %s\n", resp.Status, url)
	// var data []byte
	// fmt.Println(resp.Body)
	// data, _ = io.ReadAll(resp.Body)
	// fmt.Println(string(data))
}
