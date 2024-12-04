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

func login(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "public_navbar", "login")
}

func authenticate(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	user, err := models.GetUserByEmail((r.PostFormValue("email")))
	if err != nil{
		log.Println(err)
		http.Redirect(w, r, "/login", 302)
	}
	if user.Password == models.Encrypt(r.PostFormValue("password")){
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}

		cookie := http.Cookie{
			Name: "_cookie",
			Value: session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)

		//test1@test.com
		http.Redirect(w, r, "/",302)
	}else{
		http.Redirect(w, r, "/login",302)
	}

}