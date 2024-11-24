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

func GetTodos() (todos []Todo, err error) {
	cmd := `SELECT id, content, user_id, created_at FROM todos`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	// GetTodosはデータベースからすべてのtodoアイテムを取得します。
	// todosテーブルからid、content、user_id、created_atフィールドを選択するSQLクエリを実行します。
	// この関数はTodo構造体のスライスと、クエリの実行や行のスキャン中に発生したエラーを返します。
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Content, &todo.UserId, &todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	// rows.Close()を呼び出して、データベースのコネクションを閉じます。
	rows.Close()
	return todos, err
}

func (u *User) GetTodosByUser() (todos []Todo, err error) {
	cmd := `SELECT id, content, user_id, created_at FROM todos WHERE user_id = $1`
	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Content, &todo.UserId, &todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()
	return todos, err
}