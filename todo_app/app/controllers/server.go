package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"todo_app/app/models"
	"todo_app/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string

	// filesに格納していく
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	// エラーの時はpanic
	templates := template.Must(template.ParseFiles(files...))

	//difineでテンプレートを定義した場合、ExecuteTemplateでlayoutを明示的に指定する必要がある
	templates.ExecuteTemplate(w, "layout", data)
}

func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	// クッキーを取得
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		// クッキーがある場合
		// クッキーの値を使ってセッションを取得
		sess = models.Session{UUID: cookie.Value}
		// セッションが有効かどうかをチェック
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}

// TODO:パターンとして覚える
// 正規表現を使って、URLのパスをチェックする
var validPath = regexp.MustCompile("^/todos/(edit|update|delete)/([0-9]+)$")

// リクエストをパースする関数
func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// /todos/edit/1のようなURLをパース
		// validPath.FindStringSubmatch(r.URL.Path)で、URLのパスが正しいかどうかをチェック
		// 正しい場合、正規表現にマッチした文字列が返される
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			log.Fatalln("Invalid URL")
			http.NotFound(w, r)
			return
		}
		// 3番目の要素を数値に変換
		qi, err := strconv.Atoi(q[2])
		if err != nil {
			// log.Fatalln(err)
			http.NotFound(w, r)
			return
		}
		// ハンドラ関数を呼び出す
		fn(w, r, qi)
	})
}

// StartMainServer関数を追加
func StartMainServer() error {
	// ファイルサーバーを作成
	// http.FileServer関数は、指定されたディレクトリからファイルを提供するファイルサーバーを返す
	files := http.FileServer(http.Dir(config.Config.Static))
	// http.Handle関数を使って、"/static/"にアクセスしたときに、"/static/"以下のファイルを提供するように設定
	// StripPrefix関数を使って、"/static/"以下のファイルを提供するように設定
	// 今回はconfig.iniで設定されたstaticのディレクトリを反映できるようにしている
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// 第二引数にnilを渡すことで、デフォルトのマルチプレクサを使用してHTTPリクエストを処理する
	// マルチプレクサは、リクエストを適切なハンドラにルーティングするためのHTTPリクエストマルチプレクサです。
	// デフォルトのマルチプレクサは、404
	// 今回は、"/"にアクセスしたときにtop関数を呼び出すように設定している
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/authenticate", authenticate)
	// ログインしてるアカウントしか見れないページ
	http.HandleFunc("/", top)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	// 要求されたURLの先頭が"/todos/edit/"である場合、parseURL関数を使ってリクエストをパース
	http.HandleFunc("/todos/edit/", parseURL(todoEdit))
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))
	http.HandleFunc("/todos/delete/", parseURL(todoDelete))
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
