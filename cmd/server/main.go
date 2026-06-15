package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/axizkhan/go_postgresSQL/config"
	handler "github.com/axizkhan/go_postgresSQL/internal/handler/http"
	"github.com/axizkhan/go_postgresSQL/internal/logger"
	"github.com/axizkhan/go_postgresSQL/internal/repository/postgres"
	"github.com/axizkhan/go_postgresSQL/internal/service/user"
	"github.com/axizkhan/go_postgresSQL/pkg/database"
	validatorPkg "github.com/axizkhan/go_postgresSQL/pkg/validator"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)


func main() {
	cfg:= config.LoadConfig()

	logger.Init()

	defer logger.Sync()

	logger.Log.Info("configuration loaded",zap.String("enviorment",cfg.AppEnv),zap.String("port",cfg.Port))
	validatorPkg.Init()

	db:=database.NewPostgresConnection(cfg.DatabaseURL)
	defer db.Close(context.Background())

	logger.Log.Info("Databse Connected")
	repo := postgres.NewUserRepository(db)
	
	service := user.NewService(repo)

	userHandler := handler.NewUserHandler(service)

	logger.Log.Info(
	"user service initialized",
	zap.Any("service", service),
)

	app:=fiber.New(fiber.Config{AppName: "User DOB API"})

	app.Get("/health",func(c*fiber.Ctx)error{
		return c.JSON(fiber.Map{
			"status":"ok",
		})
	})

	app.Post("/users", userHandler.CreateUser)

app.Get("/users/:id", userHandler.GetUserById)

app.Get("/users", userHandler.ListUser)

app.Put("/users/:id", userHandler.UpdateUser)

app.Delete("/users/:id", userHandler.DeleteUser)

	quit:=make(chan os.Signal,1)

	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)

	go func(){
		address:=fmt.Sprintf(":%s",cfg.Port)
		logger.Log.Info("Starting server ",zap.String("address",address))

		logger.Log.Info("User DOB API started successfully");
		if err:=app.Listen(address); err != nil{
			logger.Log.Fatal(err.Error())
		}

	}()

	<-quit

	logger.Log.Info("Shutdown singal recieved")

	ctx,cancel:=context.WithTimeout(context.Background(),5*time.Second)

	defer cancel()

	if err:=app.ShutdownWithContext(ctx);err!=nil{
		logger.Log.Fatal("Failed to shutdown server",zap.Error(err))
	}

	logger.Log.Info("Server exited cleanly")
}