package wrand

import (
	"math/rand"
	"time"
)

type Randomizer interface {
	Intn(max int) int
}

type MathRandomizer struct{}

func (r MathRandomizer) Intn(max int) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max)
}
