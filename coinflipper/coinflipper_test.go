package coinflipper

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func headFlip() CoinSide {		
	return Head
}


// use closer to keep function call counter
func noStreakFlip() func() CoinSide {
	callCounter := 0

	return func() CoinSide {
		callCounter++
		if (callCounter % 2 == 0) {
			return Head;
		}
		return Tail
	}
}

func TestSuccess(t *testing.T) {
	assert.True(t, success(headFlip), "Return true for flips that always result with Head")
	assert.False(t, success(noStreakFlip()), "Return false for flips that return alternate results")
}	

func TestSimulate(t *testing.T) {

	r := Simulate()
	t.Logf("The result is %f", r)
	
	assert.NotZero(t, r, "The probability must be greater than zero")
}