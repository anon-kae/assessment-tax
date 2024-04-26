package main

import (
	"net/http"

	"github.com/anon-kae/assessment-tax/middleware"
	"github.com/labstack/echo/v4"

	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	e := echo.New()

	e.GET("/", middleware.AuthMiddleware(func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go Bootcamp!")
	}))
	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
