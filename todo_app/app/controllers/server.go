package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"todo_app/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string

	// filesに格納していく
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	// エラーの時はpanic
	templates:= template.Must(template.ParseFiles(files...))

	//difineでテンプレートを定義した場合、ExecuteTemplateでlayoutを明示的に指定する必要がある
	templates.ExecuteTemplate(w, "layout", data)
}

// StartMainServer関数を追加
func StartMainServer()error{
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
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	return http.ListenAndServe(":" + config.Config.Port, nil)
}