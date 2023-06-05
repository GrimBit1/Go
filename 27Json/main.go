package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"name"`
	Age      int      `json:"age"`
	Platform string   `json:"platform"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("In the Json")
	DecodeJson()
}

func EncodeJson() {
	c1 := course{"Aditya", 18, "mobile", []string{"hehe", "hehe"}}
	jsondata, err := json.MarshalIndent(c1, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsondata))
}
func DecodeJson() {
	jsonData := `{
        "name": "Aditya",
        "age": 18,
        "platform": "mobile",
        "tags": [
                "hehe",
                "hehe"
        ]
}`
	var jsontemp map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &jsontemp)
	if err != nil {
		panic(err)
	}
	fmt.Println(jsontemp)
}
