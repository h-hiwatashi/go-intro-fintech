package config

import (
	"log"
	"todo_app/utils"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port string
	SQLdriver string
	DbName string
	User string
	LogFile string
	DbPassword string
	Host      string
    DbPort    string
	Sslmode string
	Static string
}

var Config ConfigList

func init(){
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig(){
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port: cfg.Section("web").Key("port").MustString("8080"),
		SQLdriver: cfg.Section("db").Key("driver").String(),
		DbName: cfg.Section("db").Key("name").String(),
		User: cfg.Section("db").Key("user").String(),
		DbPassword: cfg.Section("db").Key("password").String(),
		Sslmode: cfg.Section("db").Key("sslmode").String(),
		Host: cfg.Section("db").Key("host").String(),
		DbPort: cfg.Section("db").Key("port").String(),
		LogFile: cfg.Section("web").Key("logfile").String(),
		Static: cfg.Section("web").Key("static").String(),
	}
}