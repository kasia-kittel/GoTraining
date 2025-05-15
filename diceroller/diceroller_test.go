package diceroller

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestSuccess(t *testing.T) {

	gen1 := func() int {
		return 1
	}

	d := Dice{
		gen: gen1,
	}

	assert.False(t, success(d, 2, 7), "Throwing two 1s should not be success")
	assert.True(t, success(d, 8, 7), "Throwing two 1s should be success")
}

func TestSimulate(t *testing.T) {

	r := Simulate()
	t.Logf("The result is %f", r)
	
	assert.NotZero(t, r, "The probability must be greater than zero")
}