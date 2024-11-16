package main

import (
	"database/sql"
	"log"

	//使用しないため_でインポート
	_ "github.com/lib/pq"
)

var Db *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	// DB名を指定してDBを開く
	// 今回はexample.sql
	//今回は非sslモードで接続
	// var err error
	Db, err := sql.Open("postgres", "user=test_user dbname=testdb password=password sslmode=disable")
	if err != nil{
		log.Panicln(err)
	}
	defer Db.Close()
}