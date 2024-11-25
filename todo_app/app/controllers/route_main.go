package controllers

import (
	"html/template"
	"log"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request){
	// ParseFiles関数を使って、テンプレートファイルをパースする
	// htmlテンプレートファイルをパースし、テンプレートオブジェクトが返される
	// テンプレートオブジェクトのExecuteメソッドを使って、テンプレートをブラウザにレンダリングする
	t, err := template.ParseFiles("app/views/templates/top.html")
	if err != nil {
		log.Fatalln(err)
	}
	// 第二引数
	// Executeメソッドの第二引数には、テンプレートに埋め込むデータを渡す
	// ここでは、文字列"Hello, World"を渡している
	// テンプレートファイルには、{{.}}という記述があり、これはExecuteメソッドの第二引数で渡したデータを埋め込む場所を示している
	err = t.Execute(w, "Hello")
	if err != nil {
        log.Fatal(err)
    }
}