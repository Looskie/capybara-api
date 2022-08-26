package v1

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"math/rand"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	. "github.com/looskie/capybara-api/utils"
)

/*
To the left [To the left
To the right [To the right
Find your ride and put that whip in drive and do the FENDI Slide
Now show me how you slide [That boy is being hunted
To the left [To the left
To the right [To the right
Find your ride, now put that whip in drive
And do the FENDI Slide, now show me how you slide

I'm outside your crib, jump inside
Heard this party jumping, we gon' slide
Straight to the front of the line
They just happy that we arrived
I'm the type to finish what I started [Yeah, yeah
If I really want it then I bought it [Yeah, yeah
She so ? when she saw me [She did
Cuz' I FENDI slide [Slide
Through the party [I slid
I don't know your man but he knows me [Ha, ha
When you gettin' money, you don't get lonely [I am never lonely
Girl, you working and I really like that [I do
You bout' to make me miss my flight back
*/

func GetCapybaras(c *fiber.Ctx) error {
	var from = c.Query("from")
	var take = c.Query("take")
	var random = c.Query("random")

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

		/* if user wants random index */
		var index = i
		if random == "true" {
			index = rand.Intn(NUMBER_OF_IMAGES-parsedFrom) + parsedFrom
		}

		file, err := os.Open("./capys/capy" + fmt.Sprint(index) + ".jpg")

		if err != nil {
			println(err.Error())
		}

		image, _, err := image.DecodeConfig(file)

		if err != nil {
			println(err.Error())
		}

		photos = append(photos, ImageStruct{
			URL:    c.BaseURL() + "/v1/capybara/" + fmt.Sprint(index),
			Index:  index,
			Width:  image.Width,
			Height: image.Height,
		})

		file.Close()
	}

	return c.JSON(Response{
		Success: true,
		Data:    photos,
	})
}
