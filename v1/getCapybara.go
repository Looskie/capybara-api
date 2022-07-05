package v1

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/gofiber/fiber/v2"
	splash "github.com/hbagdi/go-unsplash/unsplash"
	"github.com/looskie/capybara-api/utils"
)

func GetCapybara(c *fiber.Ctx) error {
	var unsplash = utils.Unsplash()
	var wantsJSON = c.Query("json")

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

	if wantsJSON == "true" {
		return c.JSON(Response{
			Success: true,
			Data:    (*randomCapys.Results)[randomIndex],
		})
	}

	bytes, err := downloadFile((*randomCapys.Results)[randomIndex].Urls.Regular.String())

	if err != nil {
		return c.JSON(Response{
			Success: false,
			Message: err.Error(),
		})
	}

	body, err := ioutil.ReadAll(bytes)

	if err != nil {
		return c.SendString("An error occurred")
	}

	return c.Send(body)
}

func downloadFile(URL string) (io.ReadCloser, error) {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Received non 200 response code")
	}

	return response.Body, nil
}
