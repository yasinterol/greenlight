package main

import (
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"os"
	"time"
)

type config struct {
	port string
	env  string
}

type application struct {
	config config
	logger *slog.Logger
}

func main() {
	var cfg config

	flag.StringVar(&cfg.port, "port", "4000", "Api Server Host")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		config: cfg,
		logger: logger,
	}

	fiberApp := fiber.New(
		fiber.Config{
			WriteTimeout: 10 * time.Second,
			ReadTimeout:  5 * time.Second,
		})

	app.routes(fiberApp)

	app.logger.Info(fmt.Sprintf("Starting server on port %d in %s mode", app.config.port, app.config.env))
	err := fiberApp.Listen(fmt.Sprintf(":%s", app.config.port))
	if err != nil {
		app.logger.Error(err.Error())
		os.Exit(1)
	}
}
