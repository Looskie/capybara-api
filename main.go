package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	v1 "github.com/looskie/capybara-api/v1"
)

func main() {
	godotenv.Load()

	app := fiber.New()
	app.Use(recover.New(recover.Config{
		Next:             nil,
		EnableStackTrace: true,
	}))

	app.Use(logger.New(logger.Config{
		Format: "${time} |   ${cyan}${status} ${reset}|   ${latency} | ${ip} on ${cyan}${ua} ${reset}| ${cyan}${method} ${reset}${path} \n",
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(v1.Response{
			Success: true,
			Message: "ok you pull up",
		})
	})

	v1Group := app.Group("/v1")
	v1Group.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(v1.Response{
			Success: true,
			Message: "welcome to v1 of capybara heaven",
		})
	})

	v1Group.Get("/capybaras", v1.GetCapybaras)
	v1Group.Get("/capybara", v1.GetCapybara)
	v1Group.Get("/capybara/:index", v1.GetCapybaraByIndex)

	app.Listen(":3000")
}
