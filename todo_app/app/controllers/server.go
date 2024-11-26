package controllers

import (
	"net/http"
	"todo_app/config"
)

// StartMainServer関数を追加
func StartMAinServer()error{
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
	return http.ListenAndServe(":" + config.Config.Port, nil)
}