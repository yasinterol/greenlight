package main

import "github.com/gofiber/fiber/v2"

func (app *application) routes(fiberApp *fiber.App) {
	router := fiberApp.Group("/v1")
	router.Get("/healthcheck", app.healtcheckHandler)
	router.Post("/movies", app.createMovieHandler)
	router.Get("/movies/:id", app.showMovieHandler)
}
