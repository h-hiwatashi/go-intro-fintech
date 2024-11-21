package models

import (
	"database/sql"
	"fmt"
	"log"
	"todo_app/config"

	_ "github.com/lib/pq"
)

var Db *sql.DB

var err error

const (
	tabeleNameUser = "users"
)

func init() {
	dataSourceName := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", config.Config.User, config.Config.DbPassword, config.Config.DbName, config.Config.Host, config.Config.DbPort, config.Config.Sslmode)
	Db, err = sql.Open(config.Config.SQLdriver, dataSourceName)
	if err != nil {
		log.Fatalln(err)
	}


	comU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
        id SERIAL PRIMARY KEY,
		uuid TEXT NOT NULL UNIQUE,
		name TEXT,
		email TEXT,
		password TEXT,
		created_at TIMESTAMP)`, tabeleNameUser)
	// Dbインスタンスを使ってSQLを実行する
    _, err = Db.Exec(comU)
    if err != nil {
        log.Fatalln(err)
    }
}