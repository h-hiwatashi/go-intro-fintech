package main

import "database/sql"

var Db *sql.DB

func dbSample() {
	Db, err := sql.Open("mysql", "root:@/test")
	if err != nil {
		panic(err)
	}
	defer Db.Close()
}

func main() {
	dbSample()
}