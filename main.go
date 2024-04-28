package main

import (
	"net/http"

	"github.com/anon-kae/assessment-tax/postgres"
	"github.com/anon-kae/assessment-tax/controllers"
	"github.com/anon-kae/assessment-tax/helper"
	"github.com/anon-kae/assessment-tax/middleware"
	"github.com/labstack/echo/v4"

	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db, err := postgres.New(postgres.Configs{DatabaseURL: os.Getenv("DATABASE_URL")})

	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Use(middleware.ErrorHandler)
	e.Validator = helper.NewValidator()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go Bootcamp!")
	})
	// e.GET("/", middleware.AuthMiddleware(func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, Go Bootcamp!")
	// }))

	taxController := controllers.New(db)
	taxController.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
