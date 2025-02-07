package main

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (app *application) readIDParam(c *fiber.Ctx) (int64, error) {
	ID := c.Params("id")
	id, err := strconv.ParseInt(ID, 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}
