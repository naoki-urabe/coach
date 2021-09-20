package config

import (
	"os"
)

type ConfigList struct {
	DbDriverName   string
	DbName         string
	DbUserName     string
	DbUserPassword string
	DbHost         string
	DbPort         string
	ApiPort        string
}

var Config ConfigList

func init() {
	Config = ConfigList{
		DbDriverName:   os.Getenv("DB_DRIVER_NAME"),
		DbName:         os.Getenv("DB_NAME"),
		DbUserName:     os.Getenv("DB_USERNAME"),
		DbUserPassword: os.Getenv("DB_USER_PASSWORD"),
		DbHost:         os.Getenv("DB_HOST"),
		DbPort:         os.Getenv("DB_PORT"),
		ApiPort:        os.Getenv("API_PORT"),
	}
}
