package v1

import (
	"fmt"
	"image"
	"math/rand"
	"os"

	"github.com/gofiber/fiber/v2"
)

func GetCapybara(c *fiber.Ctx) error {
	var wantsJSON = c.Query("json")
	randomIndex := rand.Intn(NUMBER_OF_IMAGES)

	if wantsJSON == "true" {
		file, err := os.Open("./capys/capy" + fmt.Sprint(randomIndex) + ".jpg")

		if err != nil {
			println(err.Error())
		}

		image, _, err := image.DecodeConfig(file)

		if err != nil {
			println(err.Error())
		}

		return c.JSON(Response{
			Success: true,
			Data: ImageStruct{
				URL:    c.BaseURL() + "/v1/capybara/" + fmt.Sprint(randomIndex),
				Index:  randomIndex,
				Width:  image.Width,
				Height: image.Height,
			},
		})
	}

	return c.SendFile("capys/capy" + fmt.Sprint(randomIndex) + ".jpg")
}
