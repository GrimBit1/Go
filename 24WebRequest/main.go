package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hi")
	resp, err := http.Get("https://jsonplaceholder.typicode.com/photos/20") // http is like fetch in js
	checkError(err)
	fmt.Printf("%T\n", resp)
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
	resdata, _ := io.ReadAll(resp.Body)
	fmt.Println(string(resdata))
	// http.ListenAndServe("5000",donothing)
}
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
