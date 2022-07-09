package v1

import (
	"fmt"
	"image"
	"io/ioutil"
	"math/rand"
	"os"

	"github.com/gofiber/fiber/v2"
)

func GetCapybara(c *fiber.Ctx) error {
	var wantsJSON = c.Query("json")
	randomIndex := rand.Intn(NUMBER_OF_IMAGES)

	bytes, err := ioutil.ReadFile("capys/capy" + fmt.Sprint(randomIndex) + ".jpg")

	if err != nil {
		println("error while reading capy photo", err.Error())
		if wantsJSON == "true" {
			return c.Status(500).JSON(Response{
				Success: false,
				Message: "An error occurred whilst fetching file",
			})
		}

		return c.SendStatus(500)
	}

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

	c.Set("Content-Type", "image/jpeg")
	return c.Send(bytes)
}
