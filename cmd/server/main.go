package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/axizkhan/go_postgresSQL/config"
	"github.com/axizkhan/go_postgresSQL/internal/logger"
	"github.com/axizkhan/go_postgresSQL/pkg/database"
	"github.com/gofiber/fiber/v2"
)


func main() {
	cfg:= config.LoadConfig()

	log:= logger.NewLogger()

	defer log.Sync()

	log.Info("Logger Initialized")

	db:=database.NewPostgresConnection(cfg.DatabaseURL)
	defer db.Close(context.Background())

	log.Info("Databse Connected")

	app:=fiber.New()

	app.Get("/health",func(c*fiber.Ctx)error{
		return c.JSON(fiber.Map{
			"status":"ok",
		})
	})

	quit:=make(chan os.Signal,1)

	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)

	go func(){
		address:=fmt.Sprintf(":%s",cfg.Port)
		log.Info("Starting Server on "+address)

		if err:=app.Listen(address); err != nil{
			log.Fatal(err.Error())
		}
	}()

	<-quit

	log.Info("Shutting Down Server....")

	ctx,cancel:=context.WithTimeout(context.Background(),5*time.Second)

	defer cancel()

	if err:=app.ShutdownWithContext(ctx);err!=nil{
		log.Fatal(err.Error())
	}

	log.Info("Server exited cleanly")
}