package utils

import (
	"math/rand"
)

func GetRandomIndex() int {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(NUMBER_OF_IMAGES)
	if randomIndex == 0 {
		randomIndex = 1
	}

	return randomIndex
}
