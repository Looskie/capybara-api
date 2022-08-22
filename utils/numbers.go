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

func GetRandomIndexStartingAt(num int) int {
	randomIndex := rand.Intn(NUMBER_OF_IMAGES + num)

	/* If random index is 0 or greater than random index (which is possible through the num param) */
	if randomIndex == 0 || randomIndex > NUMBER_OF_IMAGES {
		randomIndex = 1
	}

	return randomIndex
}
