package models

import (
	"log"
	"time"
)

type User struct {
	ID int
	UUID string
	Name string
	Email string
	Password string
	CreatedAt time.Time
}

// 引数なし、返り値がerrorのCreateUserメソッドを追加
func (u *User) CreateUser() (err error) {
	cmd := `INSERT INTO users (
		uuid,
		name,
		email,
		password,
		created_at) VALUES ($1, $2, $3, $4, $5)`

	// modelsパッケージで定義したDbインスタンスを使ってSQLを実行
	_, err = Db.Exec(cmd, createUUID(), u.Name, u.Email, Encrypt(u.Password), time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// 引数なし、返り値がerrorのUpdateUserメソッドを追加
func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `SELECT id, uuid, name, email, password, created_at FROM users WHERE id = $1`
	// QueryRowメソッドを使ってSQLを実行
	// Scanメソッドで取得したデータをuserに格納
	
	err = Db.QueryRow(cmd, id).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return user, err
}