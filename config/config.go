package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct{
	Port string
	AppEnv string
	DatabaseURL string
}

func LoadConfig()*Config {
	err:=godotenv.Load()

	if err!= nil{
		log.Println("No .env file found");
	}

	config:= &Config{
		Port: os.Getenv("PORT"),
		AppEnv: os.Getenv("APP_ENV"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}

	if config.Port == ""{
		log.Fatal("PORT is required")
	}

	if config.DatabaseURL==""{
		log.Fatal("DATABASE_URL is required")
	}

	return config
}