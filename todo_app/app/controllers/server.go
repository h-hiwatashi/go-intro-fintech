package controllers

import (
	"net/http"
	"todo_app/config"
)

// StartMainServer関数を追加
func StartMAinServer()error{
	// 第二引数にnilを渡すことで、デフォルトのマルチプレクサを使用してHTTPリクエストを処理する
	// マルチプレクサは、リクエストを適切なハンドラにルーティングするためのHTTPリクエストマルチプレクサです。
	// デフォルトのマルチプレクサは、404
	// 今回は、"/"にアクセスしたときにtop関数を呼び出すように設定している
	http.HandleFunc("/", top)
	return http.ListenAndServe(":" + config.Config.Port, nil)
}