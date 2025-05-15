package pi


import (
	"math/rand/v2"
)

const r int = 1000

// randomly generate points within a square and use the ratio of 
// points that fall inside an inscribed circle to estimate π.
func success(r int) bool {

	// generate 2 points inside the square
	x := rand.IntN(2*(r+1)) - r
	y := rand.IntN(2*(r+1)) - r

	// is this point also inside the circle?
	return (x*x + y*y) <= r*r
}

func Simulate() float32 {

	noOfPoints := 10000
	noOfPointsInCirle := 0

	for range noOfPoints {
		if success(r) {
			noOfPointsInCirle++
		}
	}

	return 4 * (float32(noOfPointsInCirle) / float32(noOfPoints))
}
