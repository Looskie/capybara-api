package v1

import (
	"context"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kosa3/pexels-go"
	"github.com/looskie/capybara-api/utils"
)

func GetCapybaras(c *fiber.Ctx) error {
	var page = c.Params("page")

	if len(page) == 0 {
		page = "1"
	}

	parsedPage, err := strconv.Atoi(page)

	if err != nil {
		return c.JSON(Response{
			Success: false,
			Message: err.Error(),
		})
	}

	cli := pexels.NewClient(os.Getenv("API_TOKEN"))
	ctx := context.Background()

	photos, err := utils.RedisGet(page)

	if err != nil {
		
	}

	ps, err := cli.PhotoService.Search(ctx, &pexels.PhotoParams{
		Query: "capybara",
		Page:  parsedPage,
	})

	if err != nil {
		return c.JSON(Response{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(Response{
		Success: true,
		Data:    ps.Photos,
	})
}
