package http

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestHealthRoute(t *testing.T){
	app := fiber.New()

	app.Get("/Health",func(c *fiber.Ctx)error{
		return c.SendStatus(200)
	})
	req := httptest.NewRequest("GET","/health",nil)

	res,err := app.Test(req)

	if err != nil{
		t.Fatal(err)
	}

	if res.StatusCode != 200{
		t.Errorf("expected 200, got %d",res.StatusCode)
	}
}