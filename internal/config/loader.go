package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Get() *Config{
	err:= godotenv.Load()

	if err != nil{
		log.Fatal("Error loading .env file: ", err.Error())
	}
	return &Config{
		Server: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),

		},

		Database: Database{
			Host: os.Getenv("DATABASE_HOST"),
			Port: os.Getenv("DATABASE_PORT"),
			Name: os.Getenv("DATABASE_NAME"),
			User: os.Getenv("DATABASE_USER"),
			Pass: os.Getenv("DATABASE_PASS"),
			Tz  : os.Getenv("DATABASE_TZ"),
		},
	}

}