package v1

import (
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"
	splash "github.com/hbagdi/go-unsplash/unsplash"
	"github.com/looskie/capybara-api/utils"
)

func GetCapybaras(c *fiber.Ctx) error {
	var page = c.Query("page")
	var take = c.Query("take")
	var unsplash = utils.Unsplash()

	if len(page) == 0 {
		page = "1"
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

	println(page)

	parsedPage, err := strconv.Atoi(page)

	if err != nil {
		return c.Status(500).JSON(Response{
			Success: false,
			Message: err.Error(),
		})
	}

	var photos *splash.PhotoSearchResult
	stringedPhotos, err := utils.RedisGet(page)

	if err != nil {
		fetchedPhotos, _, err := unsplash.Search.Photos(&splash.SearchOpt{
			Query:   "capybara",
			Page:    parsedPage,
			PerPage: parsedTake,
		})

		if err != nil {
			return c.Status(500).JSON(Response{
				Success: false,
				Message: "An internal server error occurred",
			})
		}

		marshalledPhotos, err := json.Marshal(fetchedPhotos)

		if err != nil {
			return c.Status(500).JSON(Response{
				Success: false,
				Message: "An internal server error occurred",
			})
		}

		utils.RedisSet(page, string(marshalledPhotos), 7)
		photos = fetchedPhotos
	} else {
		json.Unmarshal([]byte(stringedPhotos), &photos)
	}

	return c.JSON(Response{
		Success: true,
		Data:    photos,
	})
}
