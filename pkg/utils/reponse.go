package utils

import (
	"github.com/gofiber/fiber/v2"

	"github.com/axizkhan/go_postgresSQL/internal/models"
)

func Success(c *fiber.Ctx, status int, data interface{})error{
	return c.Status(status).JSON(
		models.SuccessResponse{
			Success: true,
			Data:    data,
		},
	)
}

func Error(c *fiber.Ctx, status int, message string)error{
	return c.Status(status).JSON(
		models.ErrorResponse{
			Success: false,
			Error:   message,
		},
	)
}