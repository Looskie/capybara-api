package v1

import (
	"context"
	"math/rand"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/kosa3/pexels-go"
)

func GetCapybara(c *fiber.Ctx) error {
	cli := pexels.NewClient(os.Getenv("API_TOKEN"))
	ctx := context.Background()

	min := 1
	max := 10

	page := rand.Intn(max-min) + min

	ps, err := cli.PhotoService.Search(ctx, &pexels.PhotoParams{
		Query:   "capybara",
		Page:    page,
		PerPage: 30,
	})

	if err != nil {
		return c.JSON(Response{
			Success: false,
			Message: err.Error(),
		})
	}

	if len(ps.Photos) == 0 {
		return c.JSON(Response{
			Success: true,
			Message: "An error occurred :(",
		})
	}

	index := rand.Intn(len(ps.Photos))

	println(index)

	var randomCapy = ps.Photos[index]

	return c.JSON(Response{
		Success: true,
		Data:    randomCapy,
	})
}
