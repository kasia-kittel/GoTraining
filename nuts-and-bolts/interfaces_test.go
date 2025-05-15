package nutsandbolts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestGeometry(t *testing.T) {

    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

	measureRArea, measureRPerim := Measure(r)
	measureCArea, measureCPerim := Measure(c)

	assert.Equal(t, float64(12), measureRArea)
	assert.Equal(t, float64(14), measureRPerim)

	assert.Equal(t, float64(78.53981633974483), measureCArea)
	assert.Equal(t, float64(31.41592653589793), measureCPerim)


	assert.False(t, DetectCircle(r))
	assert.True(t, DetectCircle(c))
}
