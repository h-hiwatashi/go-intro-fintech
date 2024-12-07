package controllers

import (
	"log"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	// ParseFiles関数を使って、テンプレートファイルをパースする
	// htmlテンプレートファイルをパースし、テンプレートオブジェクトが返される
	// テンプレートオブジェクトのExecuteメソッドを使って、テンプレートをブラウザにレンダリングする
	// t, err := template.ParseFiles("app/views/templates/top.html")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// 第二引数
	// Executeメソッドの第二引数には、テンプレートに埋め込むデータを渡す
	// ここでは、文字列"Hello, World"を渡している
	// テンプレートファイルには、{{.}}という記述があり、これはExecuteメソッドの第二引数で渡したデータを埋め込む場所を示している
	// err = t.Execute(w, "Hello")
	// if err != nil {
	//     log.Fatal(err)
	// }

	_, err := session(w, r)
	// sessionがなくてerrがnilでない場合、topページにリダイレクト
	if err != nil {
		generateHTML(w, "Hello", "layout", "public_navbar", "top")
	} else {
		// sessionがある場合、indexページを表示
		http.Redirect(w, r, "/todos", 302)
	}

}

// index関数を追加
func index(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	// sessionがなくてerrがnilでない場合、topページにリダイレクト
	if err != nil {
		http.Redirect(w, r, "/l", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		todos, _ := user.GetTodosByUser()
		user.Todos = todos
		// sessionがある場合、indexページを表示
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}

func todoNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "todo_new")
	}
}

func todoSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		// ParseForm関数を使って、リクエストのフォームデータを解析する
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		// セッションからユーザーを取得
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}

		// contentを取得
		content := r.PostFormValue("content")
		// ユーザーに紐づいたTodoを作成
		if err := user.CreateTodo(content); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/todos", 302)
	}
}
