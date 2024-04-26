package main

import (
	"github.com/labstack/echo/v4"
	"net/http"

	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go Bootcamp!")
	})
	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
