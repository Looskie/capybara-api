package utils

import "github.com/gofiber/fiber/v2"

func WantsJSON(c *fiber.Ctx) bool {
	if c.Query("json") == "true" {
		return true
	} else if c.Query("json") == "false" {
		return false
	}

	return string(c.Request().Header.Peek("Accept")) == "application/json"
}
