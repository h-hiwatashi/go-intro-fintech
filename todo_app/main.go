package main

import (
	"fmt"
	"todo_app/app/controllers"
	"todo_app/app/models"
)

func main() {
	fmt.Println(models.Db)
	fmt.Println("Database setup complete")

	controllers.StartMainServer()
}
