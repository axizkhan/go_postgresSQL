package database

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

func NewPostgresConnection(databaseURL string) *pgx.Conn {
	var conn *pgx.Conn
	var err error

	for i := 0; i < 10; i++ {

		conn, err = pgx.Connect(
			context.Background(),
			databaseURL,
		)

		if err == nil {
			log.Println("Connected to PostgreSQL")
			return conn
		}

		log.Printf(
			"Database not ready yet, retrying (%d/10): %v",
			i+1,
			err,
		)

		time.Sleep(3 * time.Second)
	}

	log.Fatal("Failed to connect to database:", err)

	return nil
}