package generator

import (
	"math/rand"
	"time"
)

func Generate(length, maxRand int) []int {

	rand.Seed(time.Now().UnixNano())
	a := make([]int, length)

	for i := 0; i < length; i++ { // initialize slice with random values
		a[i] = rand.Intn(maxRand)
	}
	return a
}
