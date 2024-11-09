package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	// DB名を指定してDBを開く
	// 今回はexample.sql
	Db, _ := sql.Open("sqlite3", "./example.sql")

	defer Db.Close()
	// データの追加
	//　\で複数行の文字列を書く
	// sqlインジェクションを防ぐために、?を使っている
	// cmd := "INSERT INTO persons (name, age) VALUES (?, ?)"
	// ?に入る値を引数に指定
	// _, err := Db.Exec(cmd, "Hanako", 30)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// データの更新
	// cmd := "UPDATE persons SET age = ? WHERE name = ?"
	// _, err := Db.Exec(cmd, 22, "Nancy")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// データの取得
	// cmd := "SELECT * FROM persons where age = ?"
	// row := Db.QueryRow(cmd, 22)
	// var p Person
	// err := row.Scan(&p.Name, &p.Age)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		log.Println("No row")
	// 	} else {
	// 		log.Fatalln(err)
	// 	}
	// }
	// fmt.Println(p.Name, p.Age)

	// 複数データの取得
	cmd := "SELECT * FROM persons"
	rows, _ := Db.Query(cmd)
	defer rows.Close()
	var pp []Person
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Fatalln(err)
		}
		pp = append(pp, p)
	}

	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}
}