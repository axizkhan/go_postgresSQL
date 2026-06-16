package routes

import (
	"github.com/gofiber/fiber/v2"

	handler "github.com/axizkhan/go_postgresSQL/internal/handler/http"
)

func SetupRoutes(app *fiber.App, userHandler *handler.UserHandler){
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	// User Routes
	app.Post("/users", userHandler.CreateUser)

	app.Get("/users/:id", userHandler.GetUserById)

	app.Get("/users", userHandler.ListUser)

	app.Put("/users/:id", userHandler.UpdateUser)

	app.Delete("/users/:id", userHandler.DeleteUser)

	
}