package main

import (
	"fmt"
	"log"
	"todo_app/config"
)

func main(){
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLdriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)

	log.Println("test")
}