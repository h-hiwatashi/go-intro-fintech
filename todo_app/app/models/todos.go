package models

import (
	"log"
	"time"
)

type Todo struct{
	ID int
	Content string
	UserId string
	CreatedAt time.Time
}

func (u *User) CreateTodo(content string)(err error){
	cmd := `INSERT INTO todos (
	content,
	user_id,
	created_at
	) values ($1, $2, $3)`

	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// 引数なし、返り値がerrorのUpdateUserメソッドを追加
func GetTodo(id int) (todo Todo, err error) {
	todo = Todo{}
	cmd := `SELECT id, content, user_id, created_at FROM todos WHERE id = $1`
	// QueryRowメソッドを使ってSQLを実行
	err = Db.QueryRow(cmd, id).Scan(&todo.ID, &todo.Content, &todo.UserId, &todo.CreatedAt)
	return todo, err
}