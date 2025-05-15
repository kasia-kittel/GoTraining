package diceroller

import "math/rand/v2"

type rollable interface {
	roll() int
}
 
type Dice struct {
	gen func()int
}

func (d Dice) roll() int {
	return d.gen()
}

func success(t rollable, noOfRolls int, minToSuccess int) bool {

	sum := 0

	for range noOfRolls {
		sum += t.roll()
	}

	return (sum > minToSuccess)
}

func Simulate() float32 {
	noOfTrials := 10000
	noOfSuccess := 0

	// function to simulate a dice roll - returns random value from 1 to 6
	gen := func() int {
		return rand.IntN(5) + 1
	}

	d := Dice{gen}

	for range noOfTrials {
		if success(d, 2, 7) {
			noOfSuccess++
		}
	}

	return float32(noOfSuccess) / float32(noOfTrials)
}