package main

import (
	"fmt"
	"todo_app/app/models"
)

func main(){
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLdriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	fmt.Println(models.Db)
	fmt.Println("Database setup complete")
}