package main

import (
	"database/sql"
	"log"

	//使用しないため_でインポート
	_ "github.com/lib/pq"
)

var Db *sql.DB
var err error
type Person struct {
	Name string
	Age  int
}

func main() {
	// DB名を指定してDBを開く
	//今回は非sslモードで接続
	// var err error
	Db, err = sql.Open("postgres", "host=postgres port=5432 user=root password=root dbname=dev sslmode=disable")
	// Db, err = sql.Open("postgres", "user=root password=root dbname=dev host=localhost port=5432 sslmode=disable")
	if err != nil{
		log.Panicln(err)
	}
	defer Db.Close()

	//create
	cmd := "INSERT INTO persons (name, age) VALUES ($1, $2)"
	_, err := Db.Exec(cmd, "Hanako", 30)
	if err != nil {
		log.Fatalln(err)
	}
}