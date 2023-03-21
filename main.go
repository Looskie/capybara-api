package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/looskie/capybara-api/utils"
	v1 "github.com/looskie/capybara-api/v1"
)

func main() {
	godotenv.Load()

	capyImages, _ := os.ReadDir("capys")
	utils.NUMBER_OF_IMAGES = len(capyImages)

	if err := utils.LoadCapyAlts("utils/alt.json"); err != nil {
		log.Printf("could not load alt text, using default response: %s", err)
	}

	app := fiber.New(fiber.Config{
		EnableTrustedProxyCheck: true,
		TrustedProxies:          []string{"10.50.0.0/24"},
	})

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

	app.Use(limiter.New(limiter.Config{
		Max:        500,
		Expiration: 30 * time.Second,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(429).JSON(utils.Response{
				Success: false,
				Message: "You are being rate limited",
			})
		},
		Key: func(c *fiber.Ctx) string {
			return c.GetReqHeaders()["X-Forwarded-For"]
		},
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(utils.Response{
			Success: true,
			Message: "ok you pull up (Powered by hop.io)",
		})
	})

	v1Group := app.Group("/v1")
	v1Group.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(utils.Response{
			Success: true,
			Message: "welcome to v1 of capybara heaven (Powered by hop.io)",
		})
	})

	v1Group.Get("/capybaras", v1.GetCapybaras)
	v1Group.Get("/capybara", v1.GetCapybara)
	v1Group.Get("/capybara/:index", v1.GetCapybaraByIndex)
	v1Group.Get("/capyoftheday", v1.GetCapybaraOfTheDay)
	v1Group.Get("/capyhour", v1.GetCapyHour)
	v1Group.Get("/cnrad", v1.GetConradsBirthday)

	// Capybara facts

	v1Group.Get("/fact", v1.GetCapyFact)
	v1Group.Get("/facts", v1.GetCapyFacts)

	var port = os.Getenv("PORT")

	if len(port) == 0 {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}
