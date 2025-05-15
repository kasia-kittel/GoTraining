package birthdayparadox

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func differentDays() func() int {
	callCounter := 0

	return func() int {
		callCounter++
		return callCounter
	}
}

func TestSuccess(t *testing.T) {

	day1 := func() int {
		return 1
	}

	assert.True(t, success(day1), "Having all birthdays on the same day should not be success")
	assert.False(t, success(differentDays()), "Having all birthdays on different days should not return success")
}

func TestSimulate(t *testing.T) {

	r := Simulate()
	t.Logf("The result is %f", r)
	
	assert.NotZero(t, r, "The probability must be greater than zero")
	assert.GreaterOrEqual(t, r, float32(0.48), "The probability should be close to 0.5")
}