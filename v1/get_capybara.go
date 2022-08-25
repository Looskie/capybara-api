package v1

import (
	"fmt"
	"image"
	"io/ioutil"
	"os"

	"github.com/gofiber/fiber/v2"
	. "github.com/looskie/capybara-api/utils"
)

func GetCapybara(c *fiber.Ctx) error {
	var wantsJSON = c.Query("json")
	randomIndex := GetRandomIndex()

	bytes, err := ioutil.ReadFile("./capys/capy" + fmt.Sprint(randomIndex) + ".jpg")

	c.Set("X-Capybara-Index", fmt.Sprint(randomIndex))

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

		defer file.Close()

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
