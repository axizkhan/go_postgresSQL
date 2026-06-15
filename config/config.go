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
		Port: Getenv("PORT","8080"),
		AppEnv: Getenv("APP_ENV","development"),
		DatabaseURL: RequiredEnv("DATABASE_URL"),
	}

	

	return config
}

func Getenv(key string, fallback string) string{
	value:=os.Getenv(key)
	if value==""{
		return fallback
	}
	return value
}

func RequiredEnv(key string) string{
	value:=os.Getenv(key)

	if value==""{
		log.Fatalf("%s is not in .env" ,key)
	}
	return value
}