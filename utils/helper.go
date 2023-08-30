package utils

import (
	"strconv"
	"time"
	"math/rand"
	"io/ioutil"
	"github.com/gofiber/fiber/v2"
)

func WantsJSON(c *fiber.Ctx) bool {
	if c.Query("json") == "true" {
		return true
	} else if c.Query("json") == "false" {
		return false
	}

	return string(c.Request().Header.Peek("Accept")) == "application/json"
}

/* Made our own with "://" because i dont think it should be nec. to have a
client side certificate for this project */
func BaseURL(c *fiber.Ctx) string {
	return c.Protocol() + "://" + c.Hostname()
}


// Get random seed for hourly/daily 

func SetSeed(t string) {
	var date = time.Now()
	var day = strconv.Itoa(date.Day())
	var month = strconv.Itoa(int(date.Month()))
	var year = strconv.Itoa(date.Year())
	var joined string
	if t == "hour" {
		var hour = strconv.Itoa(date.Hour())
		joined = year + month + day + hour
	} else {
		joined = year + month + day
	}
	seed, err := strconv.Atoi(joined)
	if err != nil {
		println(err.Error())
	}
	rand.Seed(int64(seed))
}


// Get index for random hour/day

func GetIndex() int {
	files, _ := ioutil.ReadDir("capys/")
	var max_rand = len(files)
	// set index
	var index = rand.Intn(max_rand) + 1
	return index
}
