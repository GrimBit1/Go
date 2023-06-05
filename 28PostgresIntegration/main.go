package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Hi")
	connStr := "user=nlab-7 dbname=nlab-7 password=nlab"
	db, err := sql.Open("postgres", connStr)

	checkErr(err)
	// age := 21
	result,err:= db.Exec("Create TABLE users(id BIGSERIAL PRIMARY KEY,name VARCHAR(50))")
	// rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
	checkErr(err)
	fmt.Println(result)
	
}
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
