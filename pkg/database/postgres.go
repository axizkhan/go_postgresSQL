package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func NewPostgresConnection(databaseURL string) *pgx.Conn{
	conn,err := pgx.Connect(context.Background(),databaseURL)

	if err != nil{
		log.Fatal("Failed to connect to database: ",err)
	}
	 log.Println("Connect to database successfully")

	 return conn
}