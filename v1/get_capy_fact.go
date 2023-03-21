package v1

import (
	"math/rand"

	"github.com/gofiber/fiber/v2"
	"github.com/looskie/capybara-api/utils"
)

func GetCapyFact(c *fiber.Ctx) error {
	factIndex := rand.Intn(len(utils.CapybaraFacts))

	return c.JSON(utils.Response{
		Success: true,
		Data: utils.FactStruct{
			Fact: utils.CapybaraFacts[factIndex],
		},
	})
}
