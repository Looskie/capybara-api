package utils

import (
	"math/rand"
)

func GetRandomIndex() int {
	randomIndex := rand.Intn(NUMBER_OF_IMAGES)
	if randomIndex == 0 {
		randomIndex = 1
	}

	return randomIndex
}
