package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"todo_app/config"

	"github.com/google/uuid"
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

// UUIDを生成する関数
func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	// SHA-1ハッシュを生成
	// SHA-1でパスワードは非推奨だが、今回は練習のために使用
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}