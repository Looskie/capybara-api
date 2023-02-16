package v1

import (
	"fmt"
	"image"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/looskie/capybara-api/utils"
)

func GetCapybaraByIndex(c *fiber.Ctx) error {
	var index = c.Params("index")
	var wantsJSON = utils.WantsJSON(c)

	parsedIndex, err := strconv.Atoi(index)
	if err != nil {
		return c.Status(500).JSON(utils.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	c.Set("X-Capybara-Index", fmt.Sprint(index))

	if wantsJSON {
		file, err := os.Open("./capys/capy" + fmt.Sprint(index) + ".jpg")

		if err != nil {
			println(err.Error())
		}

		defer file.Close()

		image, _, err := image.DecodeConfig(file)

		if err != nil {
			println(err.Error())
		}

		return c.JSON(utils.Response{
			Success: true,
			Data: utils.ImageStruct{
				URL:    utils.BaseURL(c) + "/v1/capybara/" + index,
				Index:  parsedIndex,
				Width:  image.Width,
				Height: image.Height,
				Alt:    utils.GetAlt(index),
			},
		})
	}

	return c.SendFile("capys/capy" + index + ".jpg")
}
