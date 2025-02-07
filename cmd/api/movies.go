package main

import (
	"github.com/gofiber/fiber/v2"
	"time"
	"yasintaha.erol.greenlight/internal/data"
)

func (app *application) createMovieHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"create": "true",
	})
}

func (app *application) showMovieHandler(c *fiber.Ctx) error {
	id, err := app.readIDParam(c)
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return err
	}
	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Horror IV",
		Runtime:   102,
		Genres: []string{
			"drama", "romance",
		},
		Year:    2025,
		Version: 1,
	}
	return c.Status(fiber.StatusOK).JSON(data.Envelope{"movie": movie})
}
