package main

import (
	"database/sql"
	"fmt"
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
	update(Db)
}

func create(Db *sql.DB) {
	cmd := "INSERT INTO persons (name, age) VALUES ($1, $2)"
	_, err := Db.Exec(cmd, "Tanaka", 20)
	if err != nil {
		log.Fatalln(err)
	}
}

func read(Db *sql.DB) {
	cmd := "SELECT * FROM persons where age = $1"
	// 1レコードを取得
	row := Db.QueryRow(cmd, 22)
	var p Person
	err := row.Scan(&p.Name, &p.Age)
	if err != nil {
		// データがなかったら
		if err == sql.ErrNoRows {
			log.Println("No row")
		//　それ以外のエラー
		} else {
			log.Fatalln(err)
		}
	}
	fmt.Println(p.Name, p.Age)

	// 複数データの取得(rows.Nextパターンで覚える)
	cmd = "SELECT * FROM persons"
	// 条件に合うデータを全て取得
	rows, _ := Db.Query(cmd)
	defer rows.Close()
	// Personのスライスを作成
	var pp []Person
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Fatalln(err)
		}
		// Personのスライスに追加
		pp = append(pp, p)
	}
	// 取得したデータをforループで表示
	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}
}

func update(Db *sql.DB) {
	cmd := "UPDATE persons SET age = $1 WHERE name = $2"
	_, err := Db.Exec(cmd, 22, "Hanako")
	if err != nil {
		log.Fatalln(err)
	}
}

func delete(Db *sql.DB) {
	deleteCmd := "DELETE FROM persons WHERE name = $1"
	_, err := Db.Exec(deleteCmd, "Hanako")
	if err != nil {
		log.Fatalln(err)
	}
}