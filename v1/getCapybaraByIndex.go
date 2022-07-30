package v1

import (
	"fmt"
	"image"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetCapybaraByIndex(c *fiber.Ctx) error {
	var index = c.Params("index")
	var wantsJSON = c.Query("json")

	parsedIndex, err := strconv.Atoi(index)
	if err != nil {
		return c.Status(500).JSON(Response{
			Success: false,
			Message: err.Error(),
		})
	}

	c.Set("X-Capybara-Index", fmt.Sprint(index))

	if wantsJSON == "true" {
		file, err := os.Open("./capys/capy" + fmt.Sprint(index) + ".jpg")

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
				URL:    c.BaseURL() + "/v1/capybara/" + index,
				Index:  parsedIndex,
				Width:  image.Width,
				Height: image.Height,
			},
		})
	}

	return c.SendFile("capys/capy" + index + ".jpg")
}
