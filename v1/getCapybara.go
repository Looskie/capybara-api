package v1

import (
	"encoding/json"
	"math/rand"

	"github.com/gofiber/fiber/v2"
	splash "github.com/hbagdi/go-unsplash/unsplash"
	"github.com/looskie/capybara-api/utils"
)

func GetCapybara(c *fiber.Ctx) error {
	var unsplash = utils.Unsplash()

	var randomCapys *splash.PhotoSearchResult
	stringedPhotos, err := utils.RedisGet("random")

	if err != nil {
		fetchedPhotos, _, err := unsplash.Search.Photos(&splash.SearchOpt{
			Query:   "capybara",
			PerPage: 200,
		})

		if err != nil {
			return c.JSON(Response{
				Success: false,
				Message: err.Error(),
			})
		}

		marshalledPhotos, err := json.Marshal(fetchedPhotos)

		if err != nil {
			return c.JSON(Response{
				Success: false,
				Message: "An error occurred",
			})
		}

		utils.RedisSet("random", string(marshalledPhotos), 2)
		randomCapys = fetchedPhotos
	} else {
		json.Unmarshal([]byte(stringedPhotos), &randomCapys)
	}

	randomIndex := rand.Intn(len(*randomCapys.Results))

	return c.JSON(Response{
		Success: true,
		Data:    (*randomCapys.Results)[randomIndex],
	})
}
