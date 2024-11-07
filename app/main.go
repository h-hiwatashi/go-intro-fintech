package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

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
}