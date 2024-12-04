package main

import (
	"fmt"
	"todo_app/app/models"
)

func main(){
	fmt.Println(models.Db)
	fmt.Println("Database setup complete")

	// controllers.StartMainServer()

	user, _ := models.GetUserByEmail("test1@test.com")
	fmt.Println(user)

	session, err := user.CreateSession()
	if err != nil {
		fmt.Println("Session creation failed")
	}
	fmt.Println(session)
}