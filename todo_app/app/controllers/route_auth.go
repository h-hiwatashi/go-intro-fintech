package controllers

import (
	"log"
	"net/http"
	"todo_app/app/models"
)

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		generateHTML(w, nil, "layout", "public_navbar", "signup")
	}else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}
		user := models.User{
			// signup.htmlのname属性を取得
			Name: r.PostFormValue("name"),
			// signup.htmlのemail属性を取得
			Email: r.PostFormValue("email"),
			// signup.htmlのpassword属性を取得
			Password: r.PostFormValue("password"),
		}

		if err := user.CreateUser(); err != nil {
			log.Fatalln(err)
		}
		// ユーザーが作成されたら、ログインページにリダイレクト
		http.Redirect(w, r, "/", 302)

	}
}