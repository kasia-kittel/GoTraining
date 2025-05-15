package coinflipper


import (
	"math/rand/v2"
)

type CoinSide int

const (
	Head CoinSide = iota
	Tail
)

var coinSides = map[CoinSide]string {
	Head: "head",
	Tail: "tail",
}

func (ss CoinSide) String() string {
	return coinSides[ss]
}

func randomFlip() CoinSide {
	return CoinSide(rand.IntN(2))
}

// define the success action
func success(flip func() CoinSide) bool {

	flips := make(map[CoinSide]int)

	// can be also a loop
	flips[flip()]++
	flips[flip()]++
	flips[flip()]++

	for _, v := range flips {
		if v == 3 {
			return true
		}
	}

	return false
}

func Simulate() float32 {

	count := 10000
	successCount := 0

	for range count {
		res := success(randomFlip)
		if res {
			successCount++
		}
	}

	return float32(successCount) / float32(count)
}
