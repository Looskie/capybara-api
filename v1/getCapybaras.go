package v1

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetCapybaras(c *fiber.Ctx) error {
	var from = c.Query("from")
	var take = c.Query("take")

	if len(from) == 0 {
		from = "1"
	}

	if len(take) == 0 {
		take = "25"
	}

	parsedTake, err := strconv.Atoi(take)
	if err != nil {
		return c.Status(500).JSON(Response{
			Success: false,
			Message: err.Error(),
		})
	}

	parsedFrom, err := strconv.Atoi(from)
	if err != nil {
		return c.Status(500).JSON(Response{
			Success: false,
			Message: err.Error(),
		})
	}

	var photos []ImageStruct
	for i := 0 + parsedFrom; i < parsedTake+parsedFrom && i < NUMBER_OF_IMAGES; i++ {
		file, err := os.Open("./capys/capy" + fmt.Sprint(i) + ".jpg")

		if err != nil {
			println(err.Error())
		}

		image, _, err := image.DecodeConfig(file)

		if err != nil {
			println(err.Error())
		}

		photos = append(photos, ImageStruct{
			URL:    c.BaseURL() + "/v1/capybara/" + fmt.Sprint(i),
			Index:  i,
			Width:  image.Width,
			Height: image.Height,
		})

		defer file.Close()

	}

	return c.JSON(Response{
		Success: true,
		Data:    photos,
	})
}
