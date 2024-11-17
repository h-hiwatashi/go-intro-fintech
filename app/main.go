package main

import (
	"database/sql"
	"fmt"
	"log"

	//使用しないため_でインポート
	_ "github.com/lib/pq"
	"gopkg.in/ini.v1"
)

var Db *sql.DB
var err error
type Person struct {
	Name string
	Age  int
}

// アプリケーションの状態を初期化するために用いられる関数
// パッケージ変数の初期化などに使われる
// 引数を取らず、何も値を返さない
// パッケージが初期化→パッケージ内の全ての定数と変数宣言が評価→init関数の実行
func init(){
	// iniファイルの読み込み
	cfg, _ := ini.Load("config.ini")
	// must系はデフォrト値を設定できる
	Config = ConfigList{
		Port: cfg.Section("web").Key("port").MustInt(8080),
		Dbname: cfg.Section("db").Key("dbname").MustString("example.sql"),
		SQLdriver: cfg.Section("db").Key("driver").String(),
	}
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
	// update(Db)
	fmt.Printf("Port: %d\n", Config.Port)
	fmt.Printf("Dbname: %s\n", Config.Dbname)
	fmt.Printf("SQLdriver: %s\n", Config.SQLdriver)
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

// iniファイルの説明
type ConfigList struct {
	Port int
	Dbname string
	SQLdriver string
}
var Config ConfigList