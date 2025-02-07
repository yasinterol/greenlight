package main

import (
	"github.com/gofiber/fiber/v2"
	"yasintaha.erol.greenlight/internal/data"
)

const version = "1.0.0"

type Health struct {
	Status      string
	Environment string
	Version     string
}

func (app *application) healtcheckHandler(c *fiber.Ctx) error {
	health := data.Envelope{
		"Status":      "available",
		"Environment": app.config.env,
		"Version":     version,
	}

	return c.Status(fiber.StatusOK).JSON(health)
}
