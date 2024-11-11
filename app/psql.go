package main

import (
	"database/sql"
	
	- "github.com/lib/pq"
}

var Db *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	// DB名を指定してDBを開く
	// 今回はexample.sql
	Db, _ := sql.Open("postgres", "user=postgres")
}