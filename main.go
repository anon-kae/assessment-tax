package main

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"time"

	"github.com/anon-kae/assessment-tax/controllers"
	"github.com/anon-kae/assessment-tax/helper"
	"github.com/anon-kae/assessment-tax/middleware"
	"github.com/anon-kae/assessment-tax/postgres"
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

	taxController := controllers.New(db)
	taxController.RegisterRoutes(e)

	go func() {
		if err := e.Start(os.Getenv("PORT")); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal(e.Start(os.Getenv("PORT")))
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown

	fmt.Println("shutting down the server")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
