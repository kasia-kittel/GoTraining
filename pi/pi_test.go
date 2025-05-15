package pi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimulate(t *testing.T) {

	r := Simulate()
	t.Logf("The result is %f", r)
	
	assert.NotZero(t, r, "The probability must be greater than zero")
	assert.GreaterOrEqual(t, r, float32(0.48), "The probability should be close to 0.5")
}